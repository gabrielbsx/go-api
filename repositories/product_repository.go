package repositories

import (
	"database/sql"
	"go-api/models"
	"log"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, name, price FROM products"

	rows, err := pr.connection.Query(query)

	if err != nil {
		log.Println("Error while fetching products: ", err)

		return nil, err
	}

	var products []models.Product
	var productTemp models.Product

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

func (pr *ProductRepository) CreateProduct(product models.Product) (int, error) {
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
