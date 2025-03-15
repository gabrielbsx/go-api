package main

import (
	"go-api/data"
	ProductFeature "go-api/features/products"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Data
	dbConnection, err := data.Connect()

	if err != nil {
		panic(err)
	}

	// Repository
	ProductRepository := ProductFeature.NewProductRepository(dbConnection)

	// Use cases
	ProductUsecase := ProductFeature.NewProductUsecase(ProductRepository)

	// Controllers
	ProductController := ProductFeature.NewProductController(ProductUsecase)

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
