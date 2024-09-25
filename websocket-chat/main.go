package main

import (
	"log"
	"websocket-chat/database"
	"websocket-chat/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()

	if err != nil {
		log.Fatalf("An error occured when initializing database %v", err)
	}
	defer db.Close()

	r := gin.Default()
	r.POST("/register", handlers.RegisterHandler)
	r.Run()

	// http.HandleFunc("/register", handlers.RegisterHandler)

	// port := ":8080"
    // log.Printf("Server started on port %s...", port)

	// if err := http.ListenAndServe(port, nil); err != nil {
    //     log.Fatalf("Error when starting server: %v", err)
    // }
}
