package dto

import (
	"mime/multipart"

	"github.com/sohibjon7731/nectar/internal/category/model"
)

type CategoryDTO struct {
	Title string                `form:"title" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

type CategoryProductResponseDTO struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

type CategoryResponseDTO struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

func ToCategoryResponseDTO(category model.Category) CategoryResponseDTO {

	return CategoryResponseDTO{
		ID:    category.ID,
		Title: category.Title,
		Image: category.Image,
	}
}

func ConvertToCategoryResponseDTOs(categories []model.Category) []CategoryResponseDTO {
	dtos := make([]CategoryResponseDTO, len(categories))
	for i, category := range categories {
		dtos[i] = ToCategoryResponseDTO(category)
	}
	return dtos
}
