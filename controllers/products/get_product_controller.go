package controllers

import (
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (p *ProductController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("product_id")

	if id == "" {
		response := models.Response{
			Message: "product_id is required",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := models.Response{
			Message: "product_id must be a number",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.getProductUsecase.GetProduct(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	if product == nil {
		response := models.Response{
			Message: "Product not found",
		}

		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
