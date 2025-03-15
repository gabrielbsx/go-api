package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.getProductUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, products)
}
