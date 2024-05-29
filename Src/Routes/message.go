package Routes

import "github.com/gin-gonic/gin"

func MessageRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/message", GetAllMessages)
	v1.GET("/message/:id", GetMessageById)
	v1.POST("/message", AddMessage)
	v1.PATCH("/message/:id", UpdateMessage)
	v1.DELETE("/message/:id", DeleteMessage)
}
