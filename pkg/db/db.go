package db

import (
	"github.com/aashpv/db-microservice/pkg/models"
	"github.com/jmoiron/sqlx"
)

type DataBase interface {
	CreateOrder(order models.Order) (err error)
	GetAllOrders() (orders []models.Order, err error)
	GetOrderById(id int) (order models.Order, err error)
	UpdateOrderById(order models.Order) (err error)
	DeleteOrderById(idOrder int) (err error)
	CreateProduct(product models.Product) (err error)
	GetAllProducts() (products []models.Product, err error)
	GetProductById(id int) (product models.Product, err error)
	UpdateProductById(product models.Product) (err error)
	DeleteProductById(idProduct int) (err error)
	CreateUser(user models.User) (err error)
	GetAllUsers() (users []models.User, err error)
	GetUserById(id int) (user models.User, err error)
	UpdateUserById(user models.User) (err error)
	DeleteUserById(idUser int) (err error)
}

type postgres struct {
	db *sqlx.DB
}

func New(dataSourceName string) (p DataBase, err error) {

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &postgres{db: db}, nil
}
