package router

import (
	"cuisine-api/handlers"
	"cuisine-api/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Attach middleware to log requests
	router.Use(middleware.LogRequest)

	router.HandleFunc("/api/register", handlers.Register).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(handlers.Authenticate)

	api.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
	api.HandleFunc("/recipes/{id}", handlers.GetRecipe).Methods("GET")
	api.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")
	api.HandleFunc("/recipes/{id}", handlers.UpdateRecipe).Methods("PUT")
	api.HandleFunc("/recipes/{id}", handlers.DeleteRecipe).Methods("DELETE")

	return router
}
