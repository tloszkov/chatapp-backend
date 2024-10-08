package User_Controller

import (
	"ChatApp/Src/Models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

// @Summary Get all board messages
// @ID GetAllBoardMessages
// @Tags boardmessage
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /boardmessage [get]
func GetAllBoardMessages(c *gin.Context, collection *mongo.Collection) {
	var items []bson.M
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving data from database",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &items); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading cursor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get a board message by ID
// @ID GetBoardMessageById
// @Tags boardmessage
// @Produce json
// @Param id path string true "Message Board ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /boardmessage/{id} [get]
func GetBoardMessageById(c *gin.Context, collection *mongo.Collection) {
	messageBoardIDParam := c.Param("id")
	messageBoardID, err := primitive.ObjectIDFromHex(messageBoardIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID format",
			"error":   err.Error(),
		})
		return
	}

	filter := bson.M{"_id": messageBoardID}

	var messageBoard bson.M
	if err := collection.FindOne(context.Background(), filter).Decode(&messageBoard); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving data from database",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, messageBoard)
}

// @Summary Ping the message board service
// @ID GetMessagesBoardPing
// @Tags boardmessage
// @Produce json
// @Success 200 {object} map[string]string
// @Router /boardmessage/ping [get]
func GetMessagesBoardPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// @Summary Add a new board message
// @ID AddBoardMessage
// @Tags boardmessage
// @Produce json
// @Consume json
// @Param message body Models.MessageBoard true "Message Board data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /boardmessage [post]
func AddBoardMessage(c *gin.Context, collection *mongo.Collection) {
	var newMessageBoard Models.MessageBoard
	if err := c.BindJSON(&newMessageBoard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	newMessageBoard.CreatedAt = time.Now()
	newMessageBoard.Liked = nil
	newMessageBoard.Disliked = nil
	newMessageBoard.Deleted = nil

	result, err := collection.InsertOne(context.TODO(), newMessageBoard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Board message added successfully!", "id": result.InsertedID})

}

// @Summary Update a board message by ID
// @ID UpdateBoardMessage
// @Tags boardmessage
// @Produce json
// @Consume json
// @Param id path string true "Message Board ID"
// @Param message body Models.MessageBoard true "Message Board data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /boardmessage/{id} [patch]
func UpdateBoardMessage(c *gin.Context, collection *mongo.Collection) {
	messageBoardIDParam := c.Param("id")
	messageBoardID, err := primitive.ObjectIDFromHex(messageBoardIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid message ID format",
			"error":   err.Error(),
		})
		return
	}

	var updateData Models.MessageBoard
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	updateDoc := bson.M{}
	fmt.Println(updateData)
	if updateData.MessageFrom != "" {
		updateDoc["messageFrom"] = updateData.MessageFrom
	}
	if updateData.Content != "" {
		updateDoc["content"] = updateData.Content
	}

	// fix this later (like and dislike, deleted)
	if updateData.Deleted != nil {
		updateDoc["deleted"] = 2
	}
	if updateData.Liked != nil {
		updateDoc["liked"] = 3
	}
	if updateData.Disliked != nil {
		updateDoc["disliked"] = true
	}

	update := bson.M{"$set": updateDoc}
	filter := bson.M{"_id": messageBoardID}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update board message",
			"error":   err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No board message found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Board message updated successfully!",
	})
}

// @Summary Delete a board message by ID
// @ID DeleteBoardMessageById
// @Tags boardmessage
// @Produce json
// @Param id path string true "Message Board ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /boardmessage/{id} [delete]
func DeleteBoardMessageById(c *gin.Context, collection *mongo.Collection) {
	messageBoardIDParam := c.Param("id")
	messageBoardID, err := primitive.ObjectIDFromHex(messageBoardIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID format",
			"error":   err.Error(),
		})
		return
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": messageBoardID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete board message",
			"error":   err.Error(),
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No board message found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Board message deleted successfully!",
	})
}
