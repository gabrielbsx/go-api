package controllers

import usecases "go-api/usecases/products"

type ProductController struct {
	getProductUsecase usecases.ProductUsecase
}

func NewProductController(usecase usecases.ProductUsecase) ProductController {
	return ProductController{
		getProductUsecase: usecase,
	}
}
