package usecases

import (
	"go-api/models"
)

func (u *ProductUsecase) GetProducts() ([]models.Product, error) {
	return u.repository.GetProducts()
}
