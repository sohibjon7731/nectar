package service

import (
	"errors"
	"fmt"

	"github.com/sohibjon7731/ecommerce_backend/internal/auth/model"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/repository"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/token"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/util"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/validator"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.AuthRepository
}

func NewAuthService() *AuthService {
	repo := repository.NewAuthRepository()
	return &AuthService{Repo: *repo}
}

func (s *AuthService) Register(email, password, passwordConfirmation string) (*model.User, error) {
	err := validator.EmailValidation(email)
	if err != nil {
		return nil, err
	}
	exists, err := s.Repo.ExistUserEmail(email)
	if err != nil {
		return nil, errors.New("failed to check email existence")
	}
	if exists {
		return nil, errors.New("email already taken")
	}
	err = validator.PasswordValidation(password)
	if err != nil {
		return nil, errors.New("Password is invalid")
	}

	if password != passwordConfirmation {
		return nil, errors.New("passwords do not match")
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &model.User{Email: email, Password: hashedPassword}
	err = s.Repo.CreateUser(user)
	return user, nil
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	fmt.Println(user)
	if err != nil {
		return "", "", errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid password")
	}
	accessToken, refreshToken, err := token.GenerateTokens(user.ID)
	return accessToken, refreshToken, err
}
