package db

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) CreateProduct(product models.Product) (err error) {
	_, err = p.db.Exec("INSERT INTO products (name, description, price, quantity /*image_url*/) VALUES ($1, $2, $3, $4, /*$5*/)",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		//product.ImageURL
	)

	if err != nil {
		return
	}

	return
}
