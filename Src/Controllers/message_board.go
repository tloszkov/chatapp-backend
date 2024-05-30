package Controllers

import "github.com/gin-gonic/gin"

func GetAllBoardMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All board messages!",
	})
}

func GetBoardMessageById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get board message by id!",
	})
}

func AddBoardMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Board message added!",
	})
}

func UpdateBoardMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Board message updated!",
	})
}

func DeleteBoardMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Board message deleted!",
	})
}
