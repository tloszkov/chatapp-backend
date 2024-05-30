package Controllers

import "github.com/gin-gonic/gin"

func GetAllMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All group messages!",
	})
}

func GetMessageById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get a group message by id!",
	})
}

func AddMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Group message added!",
	})
}

func UpdateMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Group message updated!",
	})
}

func DeleteMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Group message deleted!",
	})
}
