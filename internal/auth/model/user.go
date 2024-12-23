package model

import "gorm.io/gorm"

type User struct {
	Email    string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
	gorm.Model
}
