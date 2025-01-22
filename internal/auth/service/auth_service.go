package service

import (
	"errors"

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

func (s *AuthService) Register(email, password, username string) (*model.User, error) {
	if err := validator.EmailValidation(email); err != nil {
		return nil, err
	}
	
	if err:= s.checkEmailAndUsername(email, username); err != nil {
		return nil, err
	}
	
	if err := validator.PasswordValidation(password); err != nil {
		return nil, errors.New("invalid password")
	}

	

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &model.User{Email: email, Password: hashedPassword, Username: username,}
	if err:=s.Repo.CreateUser(user); err!=nil {
		return nil, errors.New("failed to create user")
	}
	return user, nil
}

func (s *AuthService) checkEmailAndUsername(email, username string)  error{
	emailExist, err:= s.Repo.ExistUserEmail(email)
	if err != nil {
		return errors.New("failed to check email existence")
	}
	if emailExist {
		return errors.New("email already taken")
	}

	usernameExist, err:= s.Repo.ExistUserUsername(username)
	if err != nil {
		return errors.New("failed to check email existence")
	}
	if usernameExist {
		return errors.New("email already taken")
	}

	return nil
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	
	if err := s.verifyPassword(user.Password, password); err != nil {
		return "", "", err
	}

	
	accessToken, refreshToken, err := token.GenerateTokens(user.ID)
	if err != nil {
		return "", "", errors.New("failed to generate tokens")
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) verifyPassword(hashedPassword, plainPassword string) error {
	
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return errors.New("invalid password")
	}
	return nil
}

