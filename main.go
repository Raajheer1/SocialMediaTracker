package main

import (
	"SocialMediaTracker/controllers"
	"SocialMediaTracker/middlewares"
	"SocialMediaTracker/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	// User routes
	protected.GET("/user", controllers.CurrentUser)

	// Social Media Account routes
	protected.POST("/accounts", controllers.RegisterAccount)

	protected.GET("/followers", controllers.GetFollowers)
	protected.POST("/followers", controllers.AddFollower)

	public.POST("/daily", controllers.Daily)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server.")
		return
	}

}
