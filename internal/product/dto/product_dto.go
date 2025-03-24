package dto

import (
	"mime/multipart"

	"github.com/sohibjon7731/nectar/internal/product/model"
)

type ProductDTO struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	CategoryID  float64               `form:"category_id" binding:"required"`
}

type ProductResponseDTO struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryID  uint    `json:"category_id"`
}

func ToProductResponseDTO(product model.Product) ProductResponseDTO {
	return ProductResponseDTO{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		CategoryID:  product.CategoryID,
	}
}

func ConvertToProductResponseDTOs(products []model.Product) []ProductResponseDTO {
	dtos := make([]ProductResponseDTO, len(products))
	for i, product := range products {
		dtos[i] = ToProductResponseDTO(product)
	}
	return dtos
}
