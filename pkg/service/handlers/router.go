package handlers

import (
	"FlowerHive/db-microservice/pkg/service/handlers/order"
	"FlowerHive/db-microservice/pkg/service/handlers/product"
	"FlowerHive/db-microservice/pkg/service/handlers/user"
	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты и обработчики для микросервиса управления товарами
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Маршрут для получения всех товаров
	r.GET("/products", product.GetAllProducts)

	// Маршрут для получения информации о конкретном товаре
	r.GET("/product/:id", product.GetProduct)

	// Маршрут для создания нового товара
	r.POST("/product/create", product.Create)

	// Маршрут для обновления информации о товаре
	r.PUT("/product/:id/update", product.Update)

	// Маршрут для удаления товара
	r.DELETE("/product/:id/delete", product.Delete)

	// Маршрут для получения всех пользователей
	r.GET("/users", user.GetAllUsers)

	// Маршрут для получения информации о конкретном пользователе
	r.GET("/user/:id", user.GetUser)

	// Маршрут для создания нового пользователя
	r.POST("/user/create", user.Create)

	// Маршрут для обновления информации о пользователе
	r.PUT("/user/:id/update", user.Update)

	// Маршрут для удаления пользователя
	r.DELETE("/user/:id/delete", user.Delete)

	// Маршрут для получения всех заказов
	r.GET("/orders", order.GetAllOrders)

	// Маршрут для получения информации о конкретном заказе
	r.GET("/order/:id", order.GetOrder)

	// Маршрут для создания нового пользователя
	r.POST("/order/create", order.Create)

	// Маршрут для обновления информации о пользователе
	r.PUT("/order/:id/update", order.Update)

	// Маршрут для удаления пользователя
	r.DELETE("/order/:id/delete", order.Delete)

	return r
}
