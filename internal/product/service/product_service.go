package service

import (
	"errors"

	"github.com/sohibjon7731/ecommerce_backend/internal/product/model"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/repository"
)

type ProductService struct{
	Repo repository.ProductRepository
}

func NewProductService() *ProductService{
	repo:= repository.NewProductRepository()
	return &ProductService{Repo: *repo}
}

func (s *ProductService) Create(image, title, description string, price float64) error{
	product:= model.Product{
		Image: image,
		Title: title,
		Description: description,
		Price: price,
	}
	err:= s.Repo.Create(&product)
	if err!=nil {
		return errors.New("Error create Product table")
	}
	return nil
}

func (s *ProductService) GetAllProducts() ([]model.Product, error){
	products, err:= s.Repo.GetAllProducts()
	if err != nil {
		return nil, errors.New("Not found products")
	}
	return products, nil
}