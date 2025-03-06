package handler

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/service"
)

type ProductHandler struct {
	Service service.ProductService
}

func NewProductHandler() *ProductHandler {
	service := service.NewProductService()
	return &ProductHandler{Service: *service}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Adds a new product with an image to the system
// @Tags         Products
// @Security	 BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        title formData string true "Product title"
// @Param        description formData string true "Product description"
// @Param        price formData number true "Product price"
// @Param        image formData file true "Product image"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /products/create [post]
func (h *ProductHandler) Create(c *gin.Context) {
	const uploadPath = "./uploads/"
	var input dto.ProductDTO
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error uploading image",
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
	err = h.Service.Create(input.Title, input.Description, input.Price, imageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "product created successfuly",
	})
}

// GetAllProducts 	godoc
// @Summary 		Get All products
// @Description 	Get All Products
// @Tags 			Products
// @Security 		BearerAuth
// @Accept 			json
// @Produce 		json
// @Success 		200 {object} dto.SuccessResponse
// @Failure 		400 {object} dto.ErrorResponse
// @Failure 		500 {object} dto.ErrorResponse
// @Router 			/products/all [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": dto.ConvertToProductResponseDTOs(products),
	})
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Updates an existing product with a new image (optional)
// @Tags         Products
// @Security	 BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id path int true "Product ID"
// @Param        title formData string false "Product title"
// @Param        description formData string false "Product description"
// @Param        price formData number false "Product price"
// @Param        image formData file false "New product image (optional)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /products/update/{id} [patch]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	fmt.Println("Received ID:", idParam)
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Product ID",
		})
		return
	}
	var productDTO dto.ProductDTO
	if err := c.ShouldBind(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	updatedProduct, err := h.Service.UpdateProduct(id, productDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, updatedProduct)
}

// DeleteProduct 	godoc
// @Summary 		delete a product
// @Description 	delete a product
// @Tags 			Products
// @Security	 	BearerAuth
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Product ID"
// @Success 		201 {object} dto.SuccessResponse
// @Failure 		400 {object} dto.ErrorResponse
// @Failure 		500 {object} dto.ErrorResponse
// @Router 			/products/delete/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid id"),
		})
		return
	}

	err = h.Service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "product deleted successfully",
	})
}
