package model

type Product struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"type:varchar(255); not null"`
	Description string `gorm:"type:text; not null"`
	Price float64 `gorm:"not null"`
	Image string `gorm:"type:varchar(255); not null"`
	
}