package products

type ProductModel struct {
	ID    int     `json:"id_product" validate:"required,gte=1"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gte=0"`
}
