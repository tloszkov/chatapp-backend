package Routes

import "github.com/gin-gonic/gin"

func GroupMessageRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/groupmessage", GetAllGroupMessages)
	v1.GET("/groupmessage/:id", GetGroupMessageById)
	v1.POST("/groupmessage", AddGroupMessage)
	v1.PATCH("/groupmessage/:id", UpdateGroupMessage)
	v1.DELETE("/groupmessage/:id", DeleteGroupMessage)
}
