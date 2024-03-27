package handlers

import (
	"github.com/aashpv/db-microservice/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handlers) GetAllProducts(c *gin.Context) {
	products, err := h.s.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve products"})
		return
	}

	// Возвращаем список товаров в формате JSON
	c.JSON(http.StatusOK, products)
}

func (h *handlers) GetProductById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// Вызываем метод GetProduct, передавая ему ID товара
	product, err := h.s.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	// Возвращаем найденный товар в формате JSON
	c.JSON(http.StatusOK, product)
}

func (h *handlers) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.s.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product created successfully"})
}

func (h *handlers) UpdateProductById(c *gin.Context) {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.s.UpdateProductById(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

func (h *handlers) DeleteProductById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = h.s.DeleteProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
