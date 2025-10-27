package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	Tags        string  `json:"tags"`
	Rating      float64 `json:"rating"`
	Reviews     int     `json:"reviews"`
}
