package db

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) CreateOrder(order models.Order) (err error) {
	_, err = p.db.Exec("INSERT INTO orders (date, status, id_product, id_user, address, other_number, other_name) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		order.Date,
		order.Status,
		order.ProductID,
		order.UserID,
		order.Address,
		order.OtherNumber,
		order.OtherName)

	if err != nil {
		return
	}

	return
}
