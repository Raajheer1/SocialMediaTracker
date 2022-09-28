package main

import (
	"SocialMediaTracker/controllers"
	"SocialMediaTracker/middlewares"
	"SocialMediaTracker/models"
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
	protected.GET("/user", controllers.CurrentUser)

	protected.GET("/followers", controllers.GetFollowers)
	protected.POST("/followers", controllers.AddFollower)

	public.POST("/daily", controllers.Daily)

	r.Run(":8080")

}
