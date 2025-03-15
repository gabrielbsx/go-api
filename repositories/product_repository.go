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
