package handlers

import (
	"github.com/aashpv/db-microservice/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	CreateOrder(c *gin.Context)
	UpdateOrderById(c *gin.Context)
	GetOrderById(c *gin.Context)
	GetAllOrders(c *gin.Context)
	DeleteOrderById(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProductById(c *gin.Context)
	GetProductById(c *gin.Context)
	GetAllProducts(c *gin.Context)
	DeleteProductById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUserById(c *gin.Context)
	GetUserById(c *gin.Context)
	GetAllUsers(c *gin.Context)
	DeleteUserById(c *gin.Context)
}

type handlers struct {
	s service.Service
}

func New(src service.Service) Handlers {
	return &handlers{s: src}
}
