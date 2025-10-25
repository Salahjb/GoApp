package handlers

import (
	"gin-backend/database"
	"gin-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load Users " + result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := database.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load User " + result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	var input models.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{
		Name:     input.Name,
		Age:      input.Age,
		Password: input.Password,
		Email:    input.Email,
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create User " + result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	var input models.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update fields if provided
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Age != 0 {
		user.Age = input.Age
	}
	if input.Password != "" {
		user.Password = input.Password
	}

	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
