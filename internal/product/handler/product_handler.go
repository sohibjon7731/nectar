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

func (h *ProductHandler) Create(c *gin.Context){
	var input dto.ProductDTO
	if err:= c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid input",
		})
		return
	}
	err:= h.Service.Create(input.Image, input.Title, input.Description, input.Price)
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

func (h *ProductHandler) GetAllProducts(c *gin.Context){
	products, err:= h.Service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products":dto.ConvertToProductDTOs(products),
	})
}
