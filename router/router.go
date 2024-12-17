package router

import (
	"github.com/gin-gonic/gin"
	"github.com/my-golang-project/controller"
)

// SetupRoutes untuk mengonfigurasi rute API
func SetupRoutes(r *gin.Engine) {
	r.GET("/users", controller.GetUsers)
	r.GET("/users/:id", controller.GetUserByID)
	r.POST("/users", controller.CreateUser)
}
