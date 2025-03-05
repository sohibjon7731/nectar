package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	service := service.NewAuthService()
	return &AuthHandler{Service: service}
}

// Register godoc
// @Summary Register a new user
// @Description This endpoint allows a new user to register with an email, password, and password confirmation.
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param register body dto.RegisterDTO true "Register User"
// @Success 200 {object} dto.Success
// @Failure 400 {object} dto.Error
// @Failure 500 {object} map[string]int
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var input dto.RegisterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Message: "invalid input",
		})
		return
	}
	access_token, refresh_token, err := h.Service.Register(input.Email, input.Password, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  access_token,
		"refresh_token": refresh_token,
	})
}

// Login godoc
// @Summary Login user
// @Description This endpoint allows a new user to login with an email, password.
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param register body dto.LoginDTO true "Login User"
// @Success 200 {object} dto.Success
// @Failure 400 {object} dto.Error
// @Failure 500 {object} map[string]int
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginDTO

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // Oshiqcha maydonlarni rad etish

	if err := decoder.Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	accessToken, refreshToken, err := h.Service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
