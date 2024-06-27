package Routes

import (
	Controllers "ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MessageBoardRoutes(router *gin.Engine, mongoClient *mongo.Client) {
	v1 := router.Group("/api/v1/")
	collection := mongoClient.Database("ChatApp").Collection("messageBoards")

	v1.GET("/boardmessage", getAllBoardMessages(collection))
	v1.GET("/boardmessage/:id", getBoardMessageById(collection))
	v1.GET("/boardmessage/ping/", getBoardMessagesPingHandler())
	v1.POST("/boardmessage", addBoardMessage(collection))
	v1.PATCH("/boardmessage/:id", updateBoardMessage(collection))
	v1.DELETE("/boardmessage/:id", deleteBoardMessage(collection))

}

func getAllBoardMessages(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetAllBoardMessages(c, collection)
	}
}

func getBoardMessageById(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetBoardMessageById(c, collection)
	}
}

func addBoardMessage(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.AddBoardMessage(c, collection)
	}
}

func updateBoardMessage(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.UpdateBoardMessage(c, collection)
	}
}

func deleteBoardMessage(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.DeleteBoardMessageById(c, collection)
	}
}

func getBoardMessagesPingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetMessagesBoardPing(c)
	}
}
