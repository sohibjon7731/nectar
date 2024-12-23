package model

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	Image string `json:"image"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}