package handlers

import (
	"github.com/aashpv/db-microservice/pkg/server/handlers/order"
	"github.com/aashpv/db-microservice/pkg/server/handlers/product"
	"github.com/aashpv/db-microservice/pkg/server/handlers/user"
	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты и обработчики для микросервиса управления товарами
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Маршрут для получения всех товаров
	r.GET("/products", product.GetAllProducts)

	// Маршрут для получения информации о конкретном товаре
	r.GET("/product/:id", product.GetProductById)

	// Маршрут для создания нового товара
	r.POST("/product/create", product.CreateProduct)

	// Маршрут для обновления информации о товаре
	r.PUT("/product/:id/update", product.UpdateProductById)

	// Маршрут для удаления товара
	r.DELETE("/product/:id/delete", product.DeleteProductById)

	// Маршрут для получения всех пользователей
	r.GET("/users", user.GetAllUsers)

	// Маршрут для получения информации о конкретном пользователе
	r.GET("/user/:id", user.GetUserById)

	// Маршрут для создания нового пользователя
	r.POST("/user/create", user.CreateUser)

	// Маршрут для обновления информации о пользователе
	r.PUT("/user/:id/update", user.UpdateUserById)

	// Маршрут для удаления пользователя
	r.DELETE("/user/:id/delete", user.DeleteUserById)

	// Маршрут для получения всех заказов
	r.GET("/orders", order.GetAllOrders)

	// Маршрут для получения информации о конкретном заказе
	r.GET("/order/:id", order.GetOrderById)

	// Маршрут для создания нового пользователя
	r.POST("/order/create", order.CreateOrder)

	// Маршрут для обновления информации о пользователе
	r.PUT("/order/:id/update", order.UpdateOrderById)

	// Маршрут для удаления пользователя
	r.DELETE("/order/:id/delete", order.DeleteOrderById)

	return r
}
