package Routes

import (
	Controllers "ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MessageRoutes(router *gin.Engine, mongoClient *mongo.Client) {
	v1 := router.Group("/api/v1/")
	collection := mongoClient.Database("ChatApp").Collection("messages")

	v1.GET("/message", getAllMessageHandler(collection))
	v1.GET("/message/:id", getMessageById(collection))
	v1.GET("/message/ping/", getMessagesPingHandler())
	v1.POST("/message", addMessageHandler(collection))
	v1.PATCH("/message/:id", updateMessageHandler(collection))
	v1.DELETE("/message/:id", deleteMessageByIdHandler(collection))
}

func getAllMessageHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetAllMessages(c, collection)
	}
}

func getMessageById(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetMessageById(c, collection)
	}
}

func addMessageHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.AddMessage(c, collection)
	}
}

func updateMessageHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.UpdateMessage(c, collection)
	}
}

func deleteMessageByIdHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.DeleteMessage(c, collection)
	}
}

func getMessagesPingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetMessagesPing(c)
	}
}
