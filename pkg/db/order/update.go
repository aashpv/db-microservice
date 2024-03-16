package order

import "FlowerHive/db-microservice/pkg/models"

func (p *postgres) Update(order models.Order) (err error) {
	_, err = p.db.Exec(
		"UPDATE orders SET status = $1 WHERE id = $2",
		order.Status,
		order.ID)

	if err != nil {
		return
	}

	return
}
