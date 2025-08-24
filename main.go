package main

import (
	"log"
	"net/http"

	"github.com/nanakwafo/go-api/config"
	"github.com/nanakwafo/go-api/controllers"
	"github.com/nanakwafo/go-api/db/sqlc"
	"github.com/nanakwafo/go-api/middleware"
	"github.com/nanakwafo/go-api/routes"

	_ "github.com/lib/pq"
)

func main() {

	db := config.ConnectDB()
	defer db.Close()

	q := sqlc.New(db)
	userController := controllers.NewUserController(q)

	router := routes.RegisterUserRoutes(userController)

	log.Println("🚀 Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", middleware.JSONMiddleware(router)))
}
