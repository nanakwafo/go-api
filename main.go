package main

import (
	"api/routes"
	"log"
	"net/http"
	"practice-go/config"
	"practice-go/middleware"
	"practice-go/routes"
)

func main() {

	db := config.ConnectDB()
	defer db.Close()

	router := routes.RegisterUserRoutes(db)

	log.Println("🚀 Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", middleware.JSONMiddleware(router)))
}
