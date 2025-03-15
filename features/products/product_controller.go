package products

import (
	shared_models "go-api/shared/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase ProductUsecase
}

func NewProductController(usecase ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product ProductModel

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	productCreated, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, productCreated)
}

func (u *ProductController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := shared_models.Response{
			Message: "ID is required",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)

	if err != nil {
		response := shared_models.Response{
			Message: "ID must be a number",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := u.productUsecase.GetProduct(productID)

	if err != nil {
		response := shared_models.Response{
			Message: err.Error(),
		}

		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := shared_models.Response{
			Message: "Product not found",
		}

		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, products)
}
