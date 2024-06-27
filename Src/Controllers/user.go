package Controllers

import (
	"ChatApp/Src/Models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetAllUsers(c *gin.Context, collection *mongo.Collection) {
	var items []bson.M
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving data from database",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &items); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading cursor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, items)
}

func GetUserById(c *gin.Context, collection *mongo.Collection) {
	userIDParam := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID format",
			"error":   err.Error(),
		})
		return
	}

	filter := bson.M{"_id": userID}

	var user bson.M
	if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving data from database",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Users API is up and running!"})
}

func AddUser(c *gin.Context, collection *mongo.Collection) {
	var newUser Models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = string(hashedPassword)

	newUser.ID = primitive.NewObjectID()
	newUser.Deleted = false
	newUser.CreatedAt = time.Now()

	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user", "error_details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added!", "id": result.InsertedID})
}

func UpdateUser(c *gin.Context, collection *mongo.Collection) {
	userIDParam := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID format",
			"error":   err.Error(),
		})
		return
	}

	var updateData Models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	update := bson.M{}

	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error hashing password",
				"error":   err.Error(),
			})
			return
		}
		updateData.Password = string(hashedPassword)
		update["password"] = updateData.Password
	}

	if updateData.FirstName != "" {
		update["firstName"] = updateData.FirstName
	}
	if updateData.LastName != "" {
		update["lastName"] = updateData.LastName
	}
	if updateData.Email != "" {
		update["email"] = updateData.Email
	}

	updateQuery := bson.M{"$set": update}
	filter := bson.M{"_id": userID}

	result, err := collection.UpdateOne(context.Background(), filter, updateQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No user found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully!",
	})
}

func DeleteUserById(c *gin.Context, collection *mongo.Collection) {
	userIDParam := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID format",
			"error":   err.Error(),
		})
		return
	}

	filter := bson.M{"_id": userID}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting user from database",
			"error":   err.Error(),
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No user found with given ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully!",
	})
}
