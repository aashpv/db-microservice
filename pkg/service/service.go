package service

import (
	"github.com/aashpv/db-microservice/pkg/db"
	"github.com/aashpv/db-microservice/pkg/models"
)

type Service interface {
	CreateOrder(body string) (err error)
	UpdateOrderById(body string) (err error)
	GetOrderById(id int) (order models.Order, err error)
	GetAllOrders() (orders []models.Order, err error)
	DeleteOrderById(id int) (err error)
	CreateProduct(body string) (err error)
	UpdateProductById(body string) (err error)
	GetProductById(id int) (product models.Product, err error)
	GetAllProducts() (products []models.Product, err error)
	DeleteProductById(id int) (err error)
	CreateUser(body string) (err error)
	UpdateUserById(body string) (err error)
	GetUserById(id int) (user models.User, err error)
	GetAllUsers() (users []models.User, err error)
	DeleteUserById(id int) (err error)
}

type service struct {
	pgs db.DataBase
}

func New(postgres db.DataBase) Service {
	return &service{pgs: postgres}
}
