package handler

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sohibjon7731/nectar/internal/category/dto"
	"github.com/sohibjon7731/nectar/internal/category/service"
)


type CategoryHandler struct {
	Service service.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	service := service.NewCategoryRepository()
	return &CategoryHandler{Service: *service}
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Adds a new category with an image to the system
// @Tags         Categories
// @Security	 BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        title formData string true "Category title"
// @Param        image formData file true "Category image"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /categories/create [post]
func (h *CategoryHandler) Create(c *gin.Context) {
	const uploadPath = "./uploads/"
	var input dto.CategoryDTO
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Image file is required or invalid",
		})
		return
	}
	savePath := filepath.Join(uploadPath, file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faylni saqlashda xatolik"})
		return
	}
	host := c.Request.Host
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	imageURL := fmt.Sprintf("%s://%s/upload/%s", scheme, host, file.Filename)
	err = h.Service.Create(input.Title, imageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "category created successfuly",
	})
}

// GetAllProducts 	godoc
// @Summary 		Get All categories
// @Description 	Get All Categories
// @Tags 			Categories
// @Security 		BearerAuth
// @Accept 			json
// @Produce 		json
// @Success 		200 {object} dto.SuccessResponse
// @Failure 		400 {object} dto.ErrorResponse
// @Failure 		500 {object} dto.ErrorResponse
// @Router 			/categories/all [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.Service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": dto.ConvertToCategoryResponseDTOs(categories),
	})
}

// UpdateCategory godoc
// @Summary      Update a category
// @Description  Updates an existing category with a new image (optional)
// @Tags         Categories
// @Security	 BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id path int true "Category ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /categories/update/{id} [patch]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	fmt.Println("Received ID:", idParam)
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Category ID",
		})
		return
	}
	var categoryDTO dto.CategoryDTO
	if err := c.ShouldBind(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	updatedCategory, err := h.Service.UpdateCategory(id, categoryDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory 	godoc
// @Summary 		delete a category
// @Description 	delete a category
// @Tags 			Categories
// @Security	 	BearerAuth
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Category ID"
// @Success 		201 {object} dto.SuccessResponse
// @Failure 		400 {object} dto.ErrorResponse
// @Failure 		500 {object} dto.ErrorResponse
// @Router 			/categories/delete/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid id"),
		})
		return
	}

	err = h.Service.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "category deleted successfully",
	})
}
