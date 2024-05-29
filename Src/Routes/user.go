package Routes

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/user", GetAllUsers)
	v1.GET("/user/:id", GetUserById)
	v1.POST("/user", AddUser)
	v1.PATCH("/user/:id", UpdateUser)
	v1.DELETE("/user/:id", DeleteUser)
}
