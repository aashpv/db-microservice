package product

import (
	"FlowerHive/db-microservice/pkg/service/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllProducts(c *gin.Context) {
	var s product.Service

	products, err := s.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve products"})
		return
	}

	// Возвращаем список товаров в формате JSON
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	var s product.Service

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// Вызываем метод GetProduct, передавая ему ID товара
	product, err := s.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	// Возвращаем найденный товар в формате JSON
	c.JSON(http.StatusOK, product)
}

func Create(c *gin.Context) {
	var body string
	var s product.Service

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := s.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product created successfully"})
}

func Update(c *gin.Context) {
	var body string
	var s product.Service

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := s.Update(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

func Delete(c *gin.Context) {
	var s product.Service

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = s.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
