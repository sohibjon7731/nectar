package model

import "github.com/sohibjon7731/nectar/internal/product/model"

type Category struct {
	ID       uint            `gorm:"primaryKey;autoIncrement"`
	Title    string          `gorm:"type:varchar(255); not null"`
	Image    string          `gorm:"type:varchar(255); not null"`
	Products []model.Product `gorm:"foreignKey:CategoryID; constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
}
