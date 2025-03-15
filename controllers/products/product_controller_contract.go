package controllers

import usecases "go-api/usecases/products"

type ProductController struct {
	productUsecase usecases.ProductUsecase
}

func NewProductController(usecase usecases.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}
