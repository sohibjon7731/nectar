package handler

import (
	"errors"
	"fmt"
	"net/http"
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
// @Summary Create a new product
// @Description Adds a new product to the system
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.ProductDTO true "Product data"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /products/create [post]
func (h *ProductHandler) Create(c *gin.Context) {
	var input dto.ProductDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}
	err := h.Service.Create(input.Title, input.Description, input.Price, input.Image)
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

// GetAllProducts godoc
// @Summary Get All products
// @Description Get All Products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /products/all [get]
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
// @Summary update a product
// @Description update a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.ProductDTO true "Product data"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /products/update/{id} [patch]
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
	if err := c.ShouldBindJSON(&productDTO); err != nil {
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


// DeleteProduct godoc
// @Summary delete a product
// @Description delete a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /products/delete/{id} [delete]
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
		"success":"product deleted successfully",
	})
}


