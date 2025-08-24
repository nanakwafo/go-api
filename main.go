package main

import (
	"log"
	"net/http"

	"github.com/nanakwafo/go-api/config"
	"github.com/nanakwafo/go-api/middleware"
	"github.com/nanakwafo/go-api/routes"
)

func main() {

	db := config.ConnectDB()
	defer db.Close()

	router := routes.RegisterUserRoutes(db)

	log.Println("🚀 Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", middleware.JSONMiddleware(router)))
}
