package handler

import "github.com/sohibjon7731/ecommerce_backend/internal/product/service"

type ProductHandler struct{
	Service service.ProductService
}

func NewProductHandler() *ProductHandler{
	service:= service.NewProductService()
	return &ProductHandler{Service: *service}
}

