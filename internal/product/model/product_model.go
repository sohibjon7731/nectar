package model

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	Image string `gorm:"type:varchar(255); not null"`
	Title string `gorm:"type:varchar(255); not null"`
	Price float64 `gorm:"not null"`
	Description string `gorm:"type:text; not null"`
	
}