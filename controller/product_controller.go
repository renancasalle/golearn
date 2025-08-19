package controller

import (
	"golearn/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {

}

func NewProductController() *ProductController {
	return &ProductController{}
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