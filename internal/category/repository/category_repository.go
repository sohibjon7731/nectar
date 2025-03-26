package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/sohibjon7731/nectar/database"
	"github.com/sohibjon7731/nectar/internal/category/dto"
	"github.com/sohibjon7731/nectar/internal/category/model"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	db, err := database.DBConnect()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	query := `INSERT INTO categories (title, image) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, &category.Title, &category.Image)
	if err != nil {
		log.Println("Error insertiong category: ", err)
		return err
	}
	return nil

}

func (r *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	query := `SELECT id, title, image FROM categories`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.Title, &c.Image)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *CategoryRepository) UpdateCategory(id uint64, updateDTO dto.CategoryDTO) (*model.Category, error) {
	query := `SELECT id, title, image FROM categories WHERE id = $1`
	row := r.DB.QueryRow(query, id)

	var category model.Category
	err := row.Scan(&category.ID, &category.Title, &category.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("category not found with this id")
		}
		return nil, err
	}

	category.Title = updateDTO.Title
	category.Image = updateDTO.Image.Filename

	updateQuery := `UPDATE categories SET title = $1, image = $2 WHERE id = $3`
	_, err = r.DB.Exec(updateQuery, category.Title, category.Image, id)
	if err != nil {
		log.Println("Error updating category:", err)
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) DeleteCategory(id uint64) error {
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM categories WHERE id=$1)`
	err := r.DB.QueryRow(checkQuery, id).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("category not found with this id")
	}

	deleteQuery := `DELETE FROM categories WHERE id=$1`
	_, err = r.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
