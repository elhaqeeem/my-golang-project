package router

import (
	"github.com/elhaqeeem/my-golang-project/controller"
	"github.com/gin-gonic/gin"
)

// SetupRoutes untuk mengonfigurasi rute API
func SetupRoutes(r *gin.Engine) {
	r.GET("/users", controller.GetUsers)
	r.GET("/users/:id", controller.GetUserByID)
	r.POST("/users", controller.CreateUser)
}
