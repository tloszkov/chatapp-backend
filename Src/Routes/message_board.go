package Routes

import "github.com/gin-gonic/gin"

func MessageBoardRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/boardmessage", GetAllBoardMessages)
	v1.GET("/boardmessage/:id", GetBoardMessageById)
	v1.POST("/boardmessage", AddBoardMessage)
	v1.PATCH("/boardmessage/:id", UpdateBoardMessage)
	v1.DELETE("/boardmessage/:id", DeleteBoardMessage)
}
