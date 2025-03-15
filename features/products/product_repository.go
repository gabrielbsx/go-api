package products

import (
	"database/sql"
	"log"
)

type ProductRepository interface {
	CreateProduct(product ProductModel) (int, error)
	GetProduct(id int) (*ProductModel, error)
	GetProducts() ([]ProductModel, error)
}

type productRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return &productRepository{
		connection: connection,
	}
}

func (pr *productRepository) GetProducts() ([]ProductModel, error) {
	query := "SELECT id, name, price FROM products"

	rows, err := pr.connection.Query(query)

	if err != nil {
		log.Println("Error while fetching products: ", err)

		return nil, err
	}

	var products []ProductModel
	var productTemp ProductModel

	for rows.Next() {
		err = rows.Scan(
			&productTemp.ID,
			&productTemp.Name,
			&productTemp.Price,
		)

		if err != nil {
			log.Println("Error while scanning products: ", err)

			return nil, err
		}

		products = append(products, productTemp)
	}

	rows.Close()

	return products, nil
}

func (pr *productRepository) CreateProduct(product ProductModel) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		log.Println("Error while preparing insert query: ", err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		log.Println("Error while inserting product: ", err)
		return 0, err
	}

	return id, nil
}

func (pr *productRepository) GetProduct(id int) (*ProductModel, error) {
	query := "SELECT id, name, price FROM products WHERE id = $1"

	row := pr.connection.QueryRow(query, id)

	var product ProductModel

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		log.Println("Error while scanning product: ", err)
		return nil, err
	}

	return &product, nil
}
