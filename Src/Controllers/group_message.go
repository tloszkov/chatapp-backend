package User_Controller

import (
	"ChatApp/Src/Models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

// @Summary Get all group messages
// @ID GetAllGroupMessages
// @Tags groupmessage
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /groupmessage [get]
func GetAllGroupMessages(c *gin.Context, collection *mongo.Collection) {
	var groupMessages []Models.GroupMessage

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving group messages",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var message Models.GroupMessage
		if err := cursor.Decode(&message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error decoding group message",
				"error":   err.Error(),
			})
			return
		}
		groupMessages = append(groupMessages, message)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cursor error while iterating group messages",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, groupMessages)
}

// @Summary Get a group message by ID
// @ID GetGroupMessageById
// @Tags groupmessage
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groupmessage/{id} [get]
func GetGroupMessageById(c *gin.Context, collection *mongo.Collection) {
	messageIDParam := c.Param("id")
	messageID, err := primitive.ObjectIDFromHex(messageIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid message ID format",
			"error":   err.Error(),
		})
		return
	}

	var message Models.GroupMessage
	filter := bson.M{"_id": messageID}

	err = collection.FindOne(context.Background(), filter).Decode(&message)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No group message found with given ID",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving group message",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, message)
}

// @Summary Ping the group message service
// @ID GetGroupMessagesPing
// @Tags groupmessage
// @Produce json
// @Success 200 {object} map[string]string
// @Router /groupmessage/ping [get]
func GetGroupMessagesPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// @Summary Add a new group message
// @ID AddGroupMessage
// @Tags groupmessage
// @Produce json
// @Consume json
// @Param message body Models.GroupMessage true "Group Message data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groupmessage [post]
func AddGroupMessage(c *gin.Context, collection *mongo.Collection) {
	var newGroupMessage Models.GroupMessage
	if err := c.ShouldBindJSON(&newGroupMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	if newGroupMessage.ID == primitive.NilObjectID {
		newGroupMessage.ID = primitive.NewObjectID()
	}
	if newGroupMessage.CreatedAt.IsZero() {
		newGroupMessage.CreatedAt = time.Now()
	}
	newGroupMessage.Liked = 0
	newGroupMessage.Disliked = 0
	newGroupMessage.Deleted = false

	result, err := collection.InsertOne(context.TODO(), newGroupMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add group message",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Group message added!",
		"id":      result.InsertedID,
	})
}

// @Summary Update a group message by ID
// @ID UpdateGroupMessage
// @Tags groupmessage
// @Produce json
// @Consume json
// @Param id path string true "Message ID"
// @Param message body Models.GroupMessage true "Group Message data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groupmessage/{id} [patch]
func UpdateGroupMessage(c *gin.Context, collection *mongo.Collection) {
	messageIDParam := c.Param("id")
	messageID, err := primitive.ObjectIDFromHex(messageIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid message ID format",
			"error":   err.Error(),
		})
		return
	}

	var updateData Models.GroupMessage
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	updateDoc := bson.M{}
	if updateData.MessageFrom != primitive.NilObjectID {
		updateDoc["messageFrom"] = updateData.MessageFrom
	}
	if updateData.Content != "" {
		updateDoc["content"] = updateData.Content
	}
	if updateData.CreatedBy != primitive.NilObjectID {
		updateDoc["createdBy"] = updateData.CreatedBy
	}
	updateDoc["liked"] = updateData.Liked
	updateDoc["disliked"] = updateData.Disliked
	updateDoc["deleted"] = updateData.Deleted

	filter := bson.M{"_id": messageID}
	update := bson.M{"$set": updateDoc}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update group message",
			"error":   err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No group message found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Group message updated successfully!",
	})
}

// @Summary Delete a group message by ID
// @ID DeleteGroupMessage
// @Tags groupmessage
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groupmessage/{id} [delete]
func DeleteGroupMessage(c *gin.Context, collection *mongo.Collection) {
	messageIDParam := c.Param("id")
	messageID, err := primitive.ObjectIDFromHex(messageIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid message ID format",
			"error":   err.Error(),
		})
		return
	}

	filter := bson.M{"_id": messageID}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete group message",
			"error":   err.Error(),
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No group message found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Group message deleted successfully!",
	})
}
