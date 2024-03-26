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

func (h *handlers) SetupRouter() *gin.Engine {
	r := gin.Default()

	// Маршрут для получения всех товаров
	r.GET("/products", h.GetAllProducts)

	// Маршрут для получения информации о конкретном товаре
	r.GET("/product/:id", h.GetProductById)

	// Маршрут для создания нового товара
	r.POST("/product/create", h.CreateProduct)

	// Маршрут для обновления информации о товаре
	r.PUT("/product/:id/update", h.UpdateProductById)

	// Маршрут для удаления товара
	r.DELETE("/product/:id/delete", h.DeleteProductById)

	// Маршрут для получения всех пользователей
	r.GET("/users", h.GetAllUsers)

	// Маршрут для получения информации о конкретном пользователе
	r.GET("/user/:id", h.GetUserById)

	// Маршрут для создания нового пользователя
	r.POST("/user/create", h.CreateUser)

	// Маршрут для обновления информации о пользователе
	r.PUT("/user/:id/update", h.UpdateUserById)

	// Маршрут для удаления пользователя
	r.DELETE("/user/:id/delete", h.DeleteUserById)

	// Маршрут для получения всех заказов
	r.GET("/orders", h.GetAllOrders)

	// Маршрут для получения информации о конкретном заказе
	r.GET("/order/:id", h.GetOrderById)

	// Маршрут для создания нового пользователя
	r.POST("/order/create", h.CreateOrder)

	// Маршрут для обновления информации о пользователе
	r.PUT("/order/:id/update", h.UpdateOrderById)

	// Маршрут для удаления пользователя
	r.DELETE("/order/:id/delete", h.DeleteOrderById)

	return r
}
