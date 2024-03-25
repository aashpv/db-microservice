package service

import (
	"encoding/json"
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

func (s *service) CreateOrder(body string) (err error) {
	var order models.Order
	err = json.Unmarshal([]byte(body), &order)
	if err != nil {
		return err
	}

	return s.pgs.CreateOrder(order)
}

func (s *service) UpdateOrderById(body string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetOrderById(id int) (order models.Order, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAllOrders() (orders []models.Order, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteOrderById(id int) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) CreateProduct(body string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateProductById(body string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetProductById(id int) (product models.Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAllProducts() (products []models.Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteProductById(id int) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) CreateUser(body string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateUserById(body string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetUserById(id int) (user models.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAllUsers() (users []models.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteUserById(id int) (err error) {
	//TODO implement me
	panic("implement me")
}

func New(postgres db.DataBase) Service {
	return &service{pgs: postgres}
}
