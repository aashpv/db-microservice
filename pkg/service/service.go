package service

import (
	"github.com/aashpv/db-microservice/pkg/db"
	"github.com/aashpv/db-microservice/pkg/models"
)

type Service interface {
	CreateOrder(order models.Order) (err error)
	UpdateOrderById(order models.Order) (err error)
	GetOrderById(id int) (order models.Order, err error)
	GetAllOrders() (orders []models.Order, err error)
	DeleteOrderById(id int) (err error)
	CreateProduct(product models.Product) (err error)
	UpdateProductById(product models.Product) (err error)
	GetProductById(id int) (product models.Product, err error)
	GetAllProducts() (products []models.Product, err error)
	DeleteProductById(id int) (err error)
	CreateUser(user models.User) (err error)
	UpdateUserById(user models.User) (err error)
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
