package Routes

import (
	"ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(router *gin.Engine, mongoClient *mongo.Client) {
	v1 := router.Group("/api/v1/")
	collection := mongoClient.Database("ChatApp").Collection("users")

	v1.GET("/user", getAllUsersHandler(collection))
	v1.GET("/user/:id", getUserByIdHandler(collection))
	v1.POST("/user", addUserHandler(collection))
	v1.PATCH("/user/:id", updateUserHandler(collection))
	v1.DELETE("/user/:id", deleteUserByIdHandler(collection))
}

func getAllUsersHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetAllUsers(c, collection)
	}
}

func getUserByIdHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.GetUserById(c, collection)
	}
}

func deleteUserByIdHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.DeleteUserById(c, collection)
	}
}

func updateUserHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.UpdateUser(c, collection)
	}
}

func addUserHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		Controllers.AddUser(c, collection)
	}
}
