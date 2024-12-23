package dto

import "github.com/sohibjon7731/ecommerce_backend/internal/product/model"


type ProductDTO struct{
	ID uint64 `json:"id"`
	Image string `json:"image"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}

func ToProductDTO(products model.Product) ProductDTO{
	return ProductDTO{
		ID: uint64(products.ID),
		Image: products.Image,
		Title: products.Title,
		Price: products.Price,
		Description: products.Description,
	}
} 

func ConvertToProductDTOs(products []model.Product) []ProductDTO {
	dtos := make([]ProductDTO, len(products))
	for i, product := range products {
		dtos[i] = ToProductDTO(product)
	}
	return dtos
}
