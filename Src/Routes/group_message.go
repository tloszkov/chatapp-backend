package Routes

import (
	Controllers "ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GroupMessageRoutes(router *gin.Engine, mongoClient *mongo.Client) {
	v1 := router.Group("/api/v1/")
	collection := mongoClient.Database("ChatApp").Collection("groupMessages")

	v1.GET("/groupmessage", getAllGroupMessageHandler(collection))
	v1.GET("/groupmessage/:id", getGroupMessageById(collection))
	v1.GET("/groupmessage/ping/", getGroupMessagesPingHandler())
	v1.POST("/groupmessage", addGroupMessage(collection))
	v1.PATCH("/groupmessage/:id", updateGroupMessage(collection))
	v1.DELETE("/groupmessage/:id", deleteGroupMessage(collection))
}

func getAllGroupMessageHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetAllGroupMessages(c, collection)
	}
}

func getGroupMessageById(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetGroupMessageById(c, collection)
	}
}

func addGroupMessage(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.AddGroupMessage(c, collection)
	}
}

func updateGroupMessage(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.UpdateGroupMessage(c, collection)
	}
}

func deleteGroupMessage(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.DeleteGroupMessage(c, collection)
	}
}

func getGroupMessagesPingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetGroupMessagesPing(c)
	}
}
