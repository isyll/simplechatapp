package main

import (
	"websocket-chat/handlers"
	"websocket-chat/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.InitDB()

	r := gin.Default()
	r.POST("/register", handlers.RegisterHandler)
	r.Run()
}
