package Controllers

import "github.com/gin-gonic/gin"

func GetAllGroupMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All messages!",
	})
}

func GetGroupMessageById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get message by id!",
	})
}

func AddGroupMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Message added!",
	})
}

func UpdateGroupMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Message updated!",
	})
}

func DeleteGroupMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Message deleted!",
	})
}
