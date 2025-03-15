package usecases

import "go-api/models"

func (u *ProductUsecase) GetProduct(id int) (*models.Product, error) {
	product, err := u.repository.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
