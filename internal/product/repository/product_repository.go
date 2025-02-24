package repository

import (
	"errors"
	"log"

	"github.com/sohibjon7731/ecommerce_backend/database"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository() *ProductRepository {
	db, err := database.DBConnect()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatal("Failed to migrate Product model")
	}

	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(id uint64, updateDTO dto.ProductDTO) (*model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	product.Title = updateDTO.Title
	product.Description = updateDTO.Description
	product.Price = updateDTO.Price
	product.Image = updateDTO.Image

	if err := r.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
