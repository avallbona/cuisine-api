package main

import (
	"cuisine-api/config"
	"cuisine-api/models"
	"cuisine-api/router"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&models.Recipe{}, &models.User{})

	r := router.Router()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
