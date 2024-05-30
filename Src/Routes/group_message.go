package Routes

import (
	"ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
)

func GroupMessageRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/groupmessage", Controllers.GetAllGroupMessages)
	v1.GET("/groupmessage/:id", Controllers.GetGroupMessageById)
	v1.POST("/groupmessage", Controllers.AddGroupMessage)
	v1.PATCH("/groupmessage/:id", Controllers.UpdateGroupMessage)
	v1.DELETE("/groupmessage/:id", Controllers.DeleteGroupMessage)
}
