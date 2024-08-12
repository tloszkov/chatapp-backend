package User_Controller

import (
	"ChatApp/Src/Models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @Summary Get all messages
// @ID GetAllMessages
// @Tags message
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /message [get]
func GetAllMessages(c *gin.Context, collection *mongo.Collection) {
	var messages []Models.Message
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving data from database",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &messages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading cursor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// @Summary Get a message by ID
// @ID GetMessageById
// @Tags message
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /message/{id} [get]
func GetMessageById(c *gin.Context, collection *mongo.Collection) {
	messageIDParam := c.Param("id")
	messageID, err := primitive.ObjectIDFromHex(messageIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid message ID format",
			"error":   err.Error(),
		})
		return
	}

	var message Models.Message
	err = collection.FindOne(context.Background(), bson.M{"_id": messageID}).Decode(&message)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No message found with given ID",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error retrieving data from database",
				"error":   err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, message)
}

// @Summary Ping the message service
// @ID GetMessagesPing
// @Tags message
// @Produce json
// @Success 200 {object} map[string]string
// @Router /message/ping [get]
func GetMessagesPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// @Summary Add a new message
// @ID AddMessage
// @Tags message
// @Produce json
// @Consume json
// @Param message body Models.Message true "Message data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /message [post]
func AddMessage(c *gin.Context, collection *mongo.Collection) {
	var newMessage Models.Message
	if err := c.ShouldBindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	if newMessage.ID == primitive.NilObjectID {
		newMessage.ID = primitive.NewObjectID()
	}
	if newMessage.CreatedAt.IsZero() {
		newMessage.CreatedAt = time.Now()
	}
	if newMessage.Seen == nil {
		defaultSeen := false
		newMessage.Seen = &defaultSeen
	}
	if newMessage.Liked == nil {
		defaultLiked := uint(0)
		newMessage.Liked = &defaultLiked
	}
	if newMessage.Disliked == nil {
		defaultDisliked := uint(0)
		newMessage.Disliked = &defaultDisliked
	}
	if newMessage.Deleted == nil {
		defaultDeleted := false
		newMessage.Deleted = &defaultDeleted
	}
	result, err := collection.InsertOne(context.TODO(), newMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add message",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message added!",
		"id":      result.InsertedID,
	})
}

// @Summary Update a message by ID
// @ID UpdateMessage
// @Tags message
// @Produce json
// @Consume json
// @Param id path string true "Message ID"
// @Param message body Models.Message true "Message data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /message/{id} [patch]
func UpdateMessage(c *gin.Context, collection *mongo.Collection) {
	messageIDParam := c.Param("id")
	messageID, err := primitive.ObjectIDFromHex(messageIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid message ID format",
			"error":   err.Error(),
		})
		return
	}

	var updateData Models.Message
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	updateDoc := bson.M{}
	if updateData.MessageFrom != "" {
		updateDoc["messageFrom"] = updateData.MessageFrom
	}
	if updateData.MessageTo != "" {
		updateDoc["messageTo"] = updateData.MessageTo
	}
	if updateData.Content != "" {
		updateDoc["content"] = updateData.Content
	}

	if updateData.Seen != nil {
		updateDoc["seen"] = *updateData.Seen
	}
	if updateData.Liked != nil {
		updateDoc["liked"] = *updateData.Liked
	}
	if updateData.Disliked != nil {
		updateDoc["disliked"] = *updateData.Disliked
	}
	if updateData.Deleted != nil {
		updateDoc["deleted"] = *updateData.Deleted
	}

	filter := bson.M{"_id": messageID}
	update := bson.M{"$set": updateDoc}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update message",
			"error":   err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No message found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message updated successfully!",
	})
}

// @Summary Delete a message by ID
// @ID DeleteMessageById
// @Tags message
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /message/{id} [delete]
func DeleteMessage(c *gin.Context, collection *mongo.Collection) {
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
			"message": "Failed to delete message",
			"error":   err.Error(),
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No message found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message deleted successfully!",
	})
}
