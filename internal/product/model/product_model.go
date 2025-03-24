package model

type Product struct {
	ID          uint
	Title       string
	Description string
	Price       float64
	Image       string
	CategoryID  uint
}
