package products

type ProductUsecase interface {
	CreateProduct(product ProductModel) (ProductModel, error)
	GetProduct(id int) (*ProductModel, error)
	GetProducts() ([]ProductModel, error)
}

type productUsecase struct {
	repository ProductRepository
}

func NewProductUsecase(repository ProductRepository) ProductUsecase {
	return &productUsecase{repository: repository}
}

func (u *productUsecase) CreateProduct(product ProductModel) (ProductModel, error) {
	productId, err := u.repository.CreateProduct(product)

	if err != nil {
		return ProductModel{}, err
	}

	product.ID = productId

	return product, nil
}

func (u *productUsecase) GetProduct(id int) (*ProductModel, error) {
	product, err := u.repository.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *productUsecase) GetProducts() ([]ProductModel, error) {
	return u.repository.GetProducts()
}
