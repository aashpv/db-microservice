package order

import (
	"FlowerHive/db-microservice/pkg/db/order"
	"FlowerHive/db-microservice/pkg/models"
)

type Service interface {
	Create(body string) (err error)
	Update(body string) (err error)
	GetOrder(id int) (order models.Order, err error)
	GetAllOrder() (orders []models.Order, err error)
	Delete(id int) (err error)
}

type service struct {
	pgs order.DataBase
}

func New(postgres order.DataBase) Service {
	return &service{pgs: postgres}
}
