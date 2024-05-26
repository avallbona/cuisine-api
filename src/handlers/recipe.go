package handlers

import (
	"cuisine-api/config"
	"cuisine-api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetRecipes handles GET requests to fetch all recipes
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var recipes []models.Recipe
	db := config.GetDB()
	db.Find(&recipes)

	json.NewEncoder(w).Encode(recipes)
}

// GetRecipe handles GET requests to fetch a specific recipe by ID
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var recipe models.Recipe
	db := config.GetDB()
	if db.First(&recipe, id).RecordNotFound() {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(recipe)
}

// CreateRecipe handles POST requests to create a new recipe
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipe models.Recipe
	_ = json.NewDecoder(r.Body).Decode(&recipe)

	db := config.GetDB()
	db.Create(&recipe)

	json.NewEncoder(w).Encode(recipe)
}

// UpdateRecipe handles PUT requests to update an existing recipe
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var recipe models.Recipe
	db := config.GetDB()
	if db.First(&recipe, id).RecordNotFound() {
		http.NotFound(w, r)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&recipe)
	recipe.ID = uint(id)
	db.Save(&recipe)

	json.NewEncoder(w).Encode(recipe)
}

// DeleteRecipe handles DELETE requests to delete a recipe by ID
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var recipe models.Recipe
	db := config.GetDB()
	if db.First(&recipe, id).RecordNotFound() {
		http.NotFound(w, r)
		return
	}

	db.Delete(&recipe)

	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}
