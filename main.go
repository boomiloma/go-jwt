package main

import (
	"go-jwt/controllers"
	"go-jwt/intializers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDb()
	intializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
