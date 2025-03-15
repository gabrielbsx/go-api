package products

type ProductUsecase struct {
	repository ProductRepository
}

func NewProductUsecase(repository ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repository}
}

func (u *ProductUsecase) CreateProduct(product ProductModel) (ProductModel, error) {
	productId, err := u.repository.CreateProduct(product)

	if err != nil {
		return ProductModel{}, err
	}

	product.ID = productId

	return product, nil
}

func (u *ProductUsecase) GetProduct(id int) (*ProductModel, error) {
	product, err := u.repository.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductUsecase) GetProducts() ([]ProductModel, error) {
	return u.repository.GetProducts()
}
