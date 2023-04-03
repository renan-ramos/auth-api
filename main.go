package main

import (
	"auth-api/controllers"
	"auth-api/initializers"
	"auth-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
