package main

import (
	"go-api/features/products"
	shared_data "go-api/shared/data"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Data
	dbConnection, err := shared_data.Connect()

	if err != nil {
		panic(err)
	}

	// Repository
	ProductRepository := products.NewProductRepository(dbConnection)

	// Use cases
	ProductUsecase := products.NewProductUsecase(ProductRepository)

	// Controllers
	ProductController := products.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products/:product_id", ProductController.GetProduct)

	server.Run(":3000")
}
