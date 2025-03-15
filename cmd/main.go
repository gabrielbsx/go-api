package main

import (
	productController "go-api/controllers/products"
	"go-api/data"
	"go-api/repositories"
	productUsecase "go-api/usecases/products"

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
	ProductRepository := repositories.NewProductRepository(dbConnection)

	// Use cases
	ProductUsecase := productUsecase.NewProductUsecase(ProductRepository)

	// Controllers
	ProductController := productController.NewProductController(ProductUsecase)

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
