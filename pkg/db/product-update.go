package db

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) UpdateProductById(product models.Product) (err error) {
	_, err = p.db.Exec(
		"UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.ID)

	if err != nil {
		return
	}

	return
}
