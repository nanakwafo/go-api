package routes

import (
	"database/sql"
	"practice-go/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUser(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser(db)).Methods("DELETE")
}
