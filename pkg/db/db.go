package db

import (
	"github.com/aashpv/db-microservice/pkg/db/order"
	"github.com/aashpv/db-microservice/pkg/db/product"
	"github.com/aashpv/db-microservice/pkg/db/user"
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

type Postgres struct {
	db      *sqlx.DB
	user    user.DataBase
	order   order.DataBase
	product product.DataBase
}

func New(
	dataSourceName string,
	user user.DataBase,
	product product.DataBase,
	order order.DataBase) (p DataBase, err error) {

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db:      db,
		user:    user,
		order:   order,
		product: product,
	}, nil
}

func (p *Postgres) UpdateOrderById(order models.Order) (err error) {
	return p.order.UpdateOrderById(order)
}

func (p *Postgres) DeleteOrderById(idOrder int) (err error) {
	return p.order.DeleteOrderById(idOrder)
}

func (p *Postgres) CreateProduct(product models.Product) (err error) {
	return p.product.CreateProduct(product)
}

func (p *Postgres) GetAllProducts() (products []models.Product, err error) {
	return p.product.GetAllProducts()
}

func (p *Postgres) GetProductById(id int) (product models.Product, err error) {
	return p.product.GetProductById(id)
}

func (p *Postgres) UpdateProductById(product models.Product) (err error) {
	return p.product.UpdateProductById(product)
}

func (p *Postgres) DeleteProductById(idProduct int) (err error) {
	return p.product.DeleteProductById(idProduct)
}

func (p *Postgres) CreateUser(user models.User) (err error) {
	return p.user.CreateUser(user)
}

func (p *Postgres) GetAllUsers() (users []models.User, err error) {
	return p.user.GetAllUsers()
}

func (p *Postgres) GetUserById(id int) (user models.User, err error) {
	return p.user.GetUserById(id)
}

func (p *Postgres) UpdateUserById(user models.User) (err error) {
	return p.user.UpdateUserById(user)
}

func (p *Postgres) DeleteUserById(idUser int) (err error) {
	return p.user.DeleteUserById(idUser)
}

func (p *Postgres) CreateOrder(order models.Order) (err error) {
	return p.order.CreateOrder(order)
}

func (p *Postgres) GetAllOrders() (orders []models.Order, err error) {
	return p.order.GetAllOrders()
}
func (p *Postgres) GetOrderById(id int) (order models.Order, err error) {
	return p.order.GetOrderById(id)
}
