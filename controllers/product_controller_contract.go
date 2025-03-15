package controllers

import "go-api/usecases"

type ProductController struct {
	getProductUsecase usecases.ProductUsecase
}

func NewProductController(usecase usecases.ProductUsecase) ProductController {
	return ProductController{
		getProductUsecase: usecase,
	}
}
