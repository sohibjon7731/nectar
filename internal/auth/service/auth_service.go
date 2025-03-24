package service

import (
	"errors"
	"fmt"

	"github.com/sohibjon7731/nectar/internal/auth/model"
	"github.com/sohibjon7731/nectar/internal/auth/repository"
	"github.com/sohibjon7731/nectar/internal/auth/token"
	"github.com/sohibjon7731/nectar/internal/auth/util"
	"github.com/sohibjon7731/nectar/internal/auth/validator"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.AuthRepository
}

func NewAuthService() *AuthService {
	repo := repository.NewAuthRepository()
	return &AuthService{Repo: *repo}
}

func (s *AuthService) Register(email, password, username string) (string, error) {
	if err := validator.EmailValidation(email); err != nil {
		return "", err
	}

	if err := s.checkEmailAndUsername(email, username); err != nil {
		return "", err
	}

	if err := validator.PasswordValidation(password); err != nil {
		return "", errors.New("invalid password")
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return "", err
	}
	user := &model.User{Email: email, Password: hashedPassword, Username: username}
	if err := s.Repo.CreateUser(user); err != nil {
		return "", errors.New("failed to create user")
	}
	fmt.Println("user", user)
	accessToken, err := token.GenerateTokens(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return accessToken, nil
}

func (s *AuthService) checkEmailAndUsername(email, username string) error {
	emailExist, err := s.Repo.ExistUserEmail(email)
	if err != nil {
		return errors.New("failed to check email existence")
	}
	if emailExist {
		return errors.New("email already taken")
	}

	usernameExist, err := s.Repo.ExistUserUsername(username)
	if err != nil {
		return errors.New("failed to check username existence")
	}
	if usernameExist {
		return errors.New("username already taken")
	}

	return nil
}

func (s *AuthService) Login(email, password string) (string, error) {

	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := s.verifyPassword(user.Password, password); err != nil {
		return "", err
	}

	accessToken,err := token.GenerateTokens(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return accessToken,nil
}

func (s *AuthService) verifyPassword(hashedPassword, plainPassword string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return errors.New("invalid password")
	}
	return nil
}
