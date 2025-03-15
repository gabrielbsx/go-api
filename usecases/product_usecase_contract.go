package usecases

import "go-api/repositories"

type ProductUsecase struct {
	repository repositories.ProductRepository
}

func NewProductUsecase(repository repositories.ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repository}
}
