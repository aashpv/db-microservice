package product

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) GetAllProducts() (products []models.Product, err error) {
	err = p.db.Select(&products, "SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	return
}

func (p *postgres) GetProductById(id int) (product models.Product, err error) {
	err = p.db.Get(&product, "SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return
	}

	return
}
