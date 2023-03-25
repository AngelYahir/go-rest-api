package main

import (
	"net/http"

	"github.com/AngelYahir/go-rest-api/db"
	"github.com/AngelYahir/go-rest-api/models"
	"github.com/AngelYahir/go-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	//? Connect to the database
	db.DBconection()
	//? Migrate the models
	db.DB.AutoMigrate(models.Task{}, models.User{})

	//? Create a new router
	route := mux.NewRouter()

	//? Index route
	route.HandleFunc("/", routes.IndexHandler)

	//? Users routes
	route.HandleFunc("/users", routes.GetUsers).Methods("GET")
	route.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	route.HandleFunc("/users", routes.CreateUser).Methods("POST")
	route.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	//? Tasks routes
	route.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	route.HandleFunc("/tasks/{id}", routes.GetTask).Methods("GET")
	route.HandleFunc("/tasks", routes.UpdTask).Methods("PUT")
	route.HandleFunc("/tasks", routes.CreateTasks).Methods("POST")
	route.HandleFunc("/tasks/{id}", routes.DeleteTasks).Methods("DELETE")

	http.ListenAndServe(":4000", route)
}
