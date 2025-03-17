package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/sohibjon7731/ecommerce_backend/database"
	"github.com/sohibjon7731/ecommerce_backend/internal/category/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/category/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	db, err := database.DBConnect()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		log.Fatal("Failed to migrate Category model")
	}
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := r.DB.Preload("Products").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) UpdateCategory(id uint64, updateDTO dto.CategoryDTO) (*model.Category, error) {
	var category model.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found with this id ")
		}
		return nil, err
	}
	category.Title = updateDTO.Title
	category.Image = updateDTO.Image.Filename

	if err := r.DB.Save(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) DeleteCategory(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&model.Category{})
	if result.Error != nil {
		fmt.Println("error_delete:", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Println("No category found with the given ID")
		return fmt.Errorf("no category found with id %d", id)
	}

	fmt.Println("Category deleted successfully")
	return nil
}