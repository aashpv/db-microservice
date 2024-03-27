package routers

import (
	"github.com/aashpv/db-microservice/pkg/server/routers/handlers"
	"github.com/gin-gonic/gin"
)

type Routers interface {
	NewRouter() *gin.Engine
}

type routers struct {
	hrs handlers.Handlers
}

func New(handlers handlers.Handlers) Routers {
	return &routers{hrs: handlers}
}

func (r *routers) NewRouter() *gin.Engine {
	engine := gin.Default()

	engine.RedirectFixedPath = true // .../about/ == .../about

	// Маршрут для получения всех товаров
	engine.GET("/products", r.hrs.GetAllProducts)

	// Маршрут для получения информации о конкретном товаре
	engine.GET("/product/:id", r.hrs.GetProductById)

	// Маршрут для создания нового товара
	engine.POST("/product/create", r.hrs.CreateProduct)

	// Маршрут для обновления информации о товаре
	engine.POST("/product/:id/update", r.hrs.UpdateProductById)

	// Маршрут для удаления товара
	engine.DELETE("/product/:id/delete", r.hrs.DeleteProductById)

	// Маршрут для получения всех пользователей
	engine.GET("/users", r.hrs.GetAllUsers)

	// Маршрут для получения информации о конкретном пользователе
	engine.GET("/user/:id", r.hrs.GetUserById)

	// Маршрут для создания нового пользователя
	engine.POST("/user/create", r.hrs.CreateUser)

	// Маршрут для обновления информации о пользователе
	engine.POST("/user/:id/update", r.hrs.UpdateUserById)

	// Маршрут для удаления пользователя
	engine.DELETE("/user/:id/delete", r.hrs.DeleteUserById)

	// Маршрут для получения всех заказов
	engine.GET("/orders", r.hrs.GetAllOrders)

	// Маршрут для получения информации о конкретном заказе
	engine.GET("/order/:id", r.hrs.GetOrderById)

	// Маршрут для создания нового пользователя
	engine.POST("/order/create", r.hrs.CreateOrder)

	// Маршрут для обновления информации о пользователе
	engine.POST("/order/:id/update", r.hrs.UpdateOrderById)

	// Маршрут для удаления пользователя
	engine.DELETE("/order/:id/delete", r.hrs.DeleteOrderById)

	//engine.Use() // auth !! оставь это для аутентификации

	return engine
}
