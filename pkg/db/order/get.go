package order

import (
	"FlowerHive/db-microservice/pkg/models"
)

func (p *postgres) GetAllOrders() (orders []models.Order, err error) {
	err = p.db.Select(&orders, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}

	return
}

func (p *postgres) GetOrderById(id int) (order models.Order, err error) {
	err = p.db.Get(&order, "SELECT * FROM orders WHERE id = $1", id)
	if err != nil {
		return
	}

	return
}
