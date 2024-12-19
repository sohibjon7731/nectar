package model

import "gorm.io/gorm"

type User struct {
	Email    string `json:"email" gorm:"unique; not null"`
	Password string `json:"password"`
	gorm.Model
}
