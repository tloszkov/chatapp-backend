package Routes

import (
	"ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/user", Controllers.GetAllUsers)
	v1.GET("/user/:id", Controllers.GetUserById)
	v1.POST("/user", Controllers.AddUser)
	v1.PATCH("/user/:id", Controllers.UpdateUser)
	v1.DELETE("/user/:id", Controllers.DeleteUser)
}
