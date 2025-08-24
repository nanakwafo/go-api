package routes

import (
	"database/sql"

	"github.com/nanakwafo/go-api/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUser(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser(db)).Methods("DELETE")
	return router
}
