package repository

import (
	"database/sql"
	"log"

	"github.com/sohibjon7731/nectar/database"
	"github.com/sohibjon7731/nectar/internal/auth/model"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository() *AuthRepository {
	db, err := database.DBConnect()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return &AuthRepository{DB: db}

}

func (r *AuthRepository) CreateUser(user *model.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("Error insertiong user: ", err)
		return err
	}
	return nil
}

func (r *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	row := r.DB.QueryRow(query, email)
	var user model.User
	var username sql.NullString

	err := row.Scan(&user.ID, &username, &user.Email, &user.Password)
	if err != nil {
		log.Println("Error fetching user: ", err)
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) ExistUserEmail(email string) (bool, error) {

	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = $1`
	err := r.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil

}

func (r *AuthRepository) ExistUserUsername(username string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = $1`
	err := r.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
