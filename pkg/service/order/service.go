package order

import (
	"github.com/aashpv/db-microservice/pkg/db/order"
	"github.com/aashpv/db-microservice/pkg/models"
)

type Service interface {
	CreateOrder(body string) (err error)
	UpdateOrderById(body string) (err error)
	GetOrderById(id int) (order models.Order, err error)
	GetAllOrders() (orders []models.Order, err error)
	DeleteOrderById(id int) (err error)
}

type service struct {
	pgs order.DataBase
}

func New(postgres order.DataBase) Service {
	return &service{pgs: postgres}
}
