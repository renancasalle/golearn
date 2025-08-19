package main

import (
	"golearn/controller"
	"golearn/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	productService := service.NewProductService()

	productController := controller.NewProductController(productService)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8080")
}