package Routes

import (
	"ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
)

func MessageBoardRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/boardmessage", Controllers.GetAllBoardMessages)
	v1.GET("/boardmessage/:id", Controllers.GetBoardMessageById)
	v1.POST("/boardmessage", Controllers.AddBoardMessage)
	v1.PATCH("/boardmessage/:id", Controllers.UpdateBoardMessage)
	v1.DELETE("/boardmessage/:id", Controllers.DeleteBoardMessage)
}
