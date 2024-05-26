package models

import "github.com/jinzhu/gorm"

type Recipe struct {
	gorm.Model
	Name         string `json:"name"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
}
