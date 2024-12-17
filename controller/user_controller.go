package controller

import (
	"net/http"

	"github.com/elhaqeeem/my-golang-project/model" // Mengimpor model.User
	"github.com/elhaqeeem/my-golang-project/service"
	"github.com/gin-gonic/gin"
)

// GetUsers untuk mengambil semua user
func GetUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID untuk mengambil user berdasarkan ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser untuk membuat user baru
func CreateUser(c *gin.Context) {
	var user model.User // Menggunakan model.User di sini
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
