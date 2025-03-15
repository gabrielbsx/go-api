package usecases

import "go-api/models"

func (u *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	productId, err := u.repository.CreateProduct(product)

	if err != nil {
		return models.Product{}, err
	}

	product.ID = productId

	return product, nil
}
