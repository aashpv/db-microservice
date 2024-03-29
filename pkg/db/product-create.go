package db

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) CreateProduct(product models.Product) (err error) {
	_, err = p.db.Exec("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
	)

	if err != nil {
		return
	}

	return
}
