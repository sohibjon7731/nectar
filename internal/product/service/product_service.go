package service

import (
	"errors"
	"fmt"

	"github.com/sohibjon7731/nectar/internal/product/dto"
	"github.com/sohibjon7731/nectar/internal/product/model"
	"github.com/sohibjon7731/nectar/internal/product/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func NewProductService() *ProductService {
	repo := repository.NewProductRepository()
	return &ProductService{Repo: *repo}
}

func (s *ProductService) Create(title, description string, price float64, image string, categoryID uint) error {
	product := model.Product{
		Title:       title,
		Description: description,
		Price:       price,
		Image:       image,
		CategoryID:  categoryID,
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

func (s *ProductService) UpdateProduct(id uint64, updateDTO dto.ProductDTO) (*model.Product, error) {
	return s.Repo.UpdateProduct(id, updateDTO)
}

func (s *ProductService) DeleteProduct(id uint64) error {
	fmt.Println("nimadir")
	return s.Repo.DeleteProduct(id)
}
