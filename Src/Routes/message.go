package Routes

import (
	"ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
)

func MessageRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/message", Controllers.GetAllMessages)
	v1.GET("/message/:id", Controllers.GetMessageById)
	v1.POST("/message", Controllers.AddMessage)
	v1.PATCH("/message/:id", Controllers.UpdateMessage)
	v1.DELETE("/message/:id", Controllers.DeleteMessage)
}
