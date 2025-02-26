package service

import (
	"errors"

	"github.com/sohibjon7731/ecommerce_backend/internal/product/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/model"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func NewProductService() *ProductService {
	repo := repository.NewProductRepository()
	return &ProductService{Repo: *repo}
}

func (s *ProductService) Create(title, description string, price float64, image string) error {
	product := model.Product{
		Title:       title,
		Description: description,
		Price:       price,
		Image:       image,
	}
	err := s.Repo.Create(&product)
	if err != nil {
		return errors.New("error create Product table")
	}
	return nil
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	products, err := s.Repo.GetAllProducts()
	if err != nil {
		return nil, errors.New("not found products")
	}
	return products, nil
}

func (s *ProductService) UpdateProduct(id uint64,updateDTO dto.ProductDTO) (*model.Product , error){
	return s.Repo.UpdateProduct(id, updateDTO)
}

func (s *ProductService) DeleteProduct(id uint64) error  {
	return s.Repo.DeleteProduct(id)
}

