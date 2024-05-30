package Controllers

import "github.com/gin-gonic/gin"

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All users!",
	})
}

func GetUserById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get user by id!",
	})
}

func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User added!",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User updated!",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User deleted!",
	})
}
