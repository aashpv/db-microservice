package handlers

import (
	"github.com/aashpv/db-microservice/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handlers) GetAllOrders(c *gin.Context) {
	orders, err := h.s.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve orders"})
		return
	}

	// Возвращаем список товаров в формате JSON
	c.JSON(http.StatusOK, orders)
}

func (h *handlers) GetOrderById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := h.s.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	// Возвращаем найденный товар в формате JSON
	c.JSON(http.StatusOK, order)
}

func (h *handlers) CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.Bind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.s.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
}

func (h *handlers) UpdateOrderById(c *gin.Context) {
	var order models.Order

	if err := c.Bind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	order.ID = idInt

	err = h.s.UpdateOrderById(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order updated successfully"})
}

func (h *handlers) DeleteOrderById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	err = h.s.DeleteOrderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
