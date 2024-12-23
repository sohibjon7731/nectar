package model

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	Image string `gorm:"type:varchar(255); not null"`
	Title string `gorm:"type:varchar(255); not null"`
	Price float64 `gorm:"type:text; not null"`
	Description string `gorm:"not null"`
}