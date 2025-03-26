package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/sohibjon7731/nectar/database"
	"github.com/sohibjon7731/nectar/internal/product/dto"
	"github.com/sohibjon7731/nectar/internal/product/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository() *ProductRepository {
	db, err := database.DBConnect()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(product *model.Product) error {
	query := `INSERT INTO products (title, description, price, image, category_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, product.Title, product.Description, product.Price, product.Image, product.CategoryID)
	return err
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	query := `SELECT id, title, description, price, image, category_id FROM products`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Image, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(id uint64, updateDTO dto.ProductDTO) (*model.Product, error) {
	var product model.Product
	query := `SELECT id, title, description, price, image, category_id FROM products WHERE id=$1`
	err := r.DB.QueryRow(query, id).Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Image, &product.CategoryID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found with this id")
		}
		return nil, err
	}

	updateQuery := `UPDATE products SET title=$1, description=$2, price=$3, image=$4 WHERE id=$5`
	_, err = r.DB.Exec(updateQuery, updateDTO.Title, updateDTO.Description, updateDTO.Price, updateDTO.Image.Filename, id)
	if err != nil {
		return nil, err
	}

	product.Title = updateDTO.Title
	product.Description = updateDTO.Description
	product.Price = updateDTO.Price
	product.Image = updateDTO.Image.Filename
	return &product, nil
}

func (r *ProductRepository) DeleteProduct(id uint64) error {
	query := `DELETE FROM products WHERE id=$1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		fmt.Println("error_delete:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No product found with the given ID")
		return fmt.Errorf("no product found with id %d", id)
	}

	fmt.Println("Product deleted successfully")
	return nil
}
