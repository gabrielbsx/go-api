package usecase

import (
	"go-api/models"
	"go-api/repositories"
)

type GetProductUsecase struct {
	repository repositories.ProductRepository
}

func NewGetProductUsecase(repository repositories.ProductRepository) GetProductUsecase {
	return GetProductUsecase{repository: repository}
}

func (u *GetProductUsecase) GetProducts() ([]models.Product, error) {
	return u.repository.GetProducts()
}
