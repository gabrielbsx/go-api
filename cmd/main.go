package main

import (
	"go-api/controllers"
	"go-api/data"
	"go-api/repositories"
	usecase "go-api/usecases"

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
	GetProductUsecase := usecase.NewGetProductUsecase(ProductRepository)

	// Controllers
	ProductController := controllers.NewProductController(GetProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":3000")
}
