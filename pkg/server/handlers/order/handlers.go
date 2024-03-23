package order

import (
	"github.com/aashpv/db-microservice/pkg/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllOrders(c *gin.Context) {
	var s order.Service

	orders, err := s.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve orders"})
		return
	}

	// Возвращаем список товаров в формате JSON
	c.JSON(http.StatusOK, orders)
}

func GetOrderById(c *gin.Context) {
	var s order.Service

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := s.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	// Возвращаем найденный товар в формате JSON
	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	var body string
	var s order.Service

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := s.CreateOrder(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
}

func UpdateOrderById(c *gin.Context) {
	var body string
	var s order.Service

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := s.UpdateOrderById(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order updated successfully"})
}

func DeleteOrderById(c *gin.Context) {
	var s order.Service

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	err = s.DeleteOrderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
