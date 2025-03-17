package service

import (
	"errors"
	"fmt"

	"github.com/sohibjon7731/ecommerce_backend/internal/category/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/category/model"
	"github.com/sohibjon7731/ecommerce_backend/internal/category/repository"
)

type CategoryService struct {
	Repo repository.CategoryRepository
}

func NewCategoryRepository() *CategoryService {
	repo := repository.NewCategoryRepository()
	return &CategoryService{
		Repo: *repo,
	}
}

func (s *CategoryService) Create(title, image string) error {
	category := model.Category{
		Title: title,
		Image: image,
	}
	err := s.Repo.Create(&category)
	if err != nil {
		return errors.New("error create Category table")
	}
	return nil
}

func (s *CategoryService) GetAllCategories() ([]model.Category, error) {
	categories, err := s.Repo.GetAllCategories()
	if err != nil {
		return nil, errors.New("not found products")
	}
	return categories, nil
}

func (s *CategoryService) UpdateCategory(id uint64, updateDTO dto.CategoryDTO) (*model.Category, error) {
	return s.Repo.UpdateCategory(id, updateDTO)
}

func (s *CategoryService) DeleteCategory(id uint64) error {
	fmt.Println("nimadir")
	return s.Repo.DeleteCategory(id)
}
