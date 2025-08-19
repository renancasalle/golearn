package controller

import (
	"golearn/model"
	"golearn/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{
		ProductService: service,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "Mouse",
			Price: 29,
		},
	}

	ctx.JSON(http.StatusOK, products)
}