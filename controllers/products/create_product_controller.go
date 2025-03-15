package controllers

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	productCreated, err := p.getProductUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, productCreated)
}
