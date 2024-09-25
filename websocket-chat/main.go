package main

import (
	"websocket-chat/handlers"
	"websocket-chat/initializers"
	"websocket-chat/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.InitDB()

	r := gin.Default()
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/test", middleware.RequireAuth, handlers.ProtectedRoute)
	r.Run()
}
