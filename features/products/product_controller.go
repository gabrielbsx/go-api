package products

import (
	shared_models "go-api/shared/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
}

type productController struct {
	usecase ProductUsecase
}

func NewProductController(usecase ProductUsecase) ProductController {
	return &productController{
		usecase: usecase,
	}
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product ProductModel

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productCreated, err := p.usecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, productCreated)
}

func (u *productController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("product_id")

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

	if productID <= 0 {
		response := shared_models.Response{
			Message: "ID must be greater than 0",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := u.usecase.GetProduct(productID)

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

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.usecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, products)
}
