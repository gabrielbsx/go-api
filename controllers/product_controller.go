package controllers

import (
	usecase "go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	getProductUsecase usecase.GetProductUsecase
}

func NewProductController(usecase usecase.GetProductUsecase) productController {
	return productController{
		getProductUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.getProductUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, products)
}
