package model

import "gorm.io/gorm"

type User struct {
	Email    string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
	Username string `gorm:"unique not null"`
	gorm.Model
}
