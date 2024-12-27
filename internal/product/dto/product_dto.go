package dto

import "github.com/sohibjon7731/ecommerce_backend/internal/product/model"

type ProductCreateDTO struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Image       string  `json:"image" binding:"required"`
}


type ProductResponseDTO struct {
	ID          uint  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}


func ToProductResponseDTO(product model.Product) ProductResponseDTO {
	return ProductResponseDTO{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
	}
}


func ConvertToProductResponseDTOs(products []model.Product) []ProductResponseDTO {
	dtos := make([]ProductResponseDTO, len(products))
	for i, product := range products {
		dtos[i] = ToProductResponseDTO(product)
	}
	return dtos
}
