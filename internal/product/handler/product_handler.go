package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/product/service"
)

type ProductHandler struct{
	Service service.ProductService
}

func NewProductHandler() *ProductHandler{
	service:= service.NewProductService()
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
func (h *ProductHandler) Create(c *gin.Context){
	var input dto.ProductDTO
	if err:= c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid input",
		})
		return
	}
	err:= h.Service.Create(input.Title, input.Description, input.Price, input.Image, )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":"product created successfuly",
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
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context){
	products, err:= h.Service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products":dto.ConvertToProductResponseDTOs(products),
	})
}
