package user

import (
	"FlowerHive/db-microservice/pkg/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	var s user.Service

	users, err := s.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}

	// Возвращаем список товаров в формате JSON
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var s user.Service

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// Вызываем метод GetProduct, передавая ему ID товара
	user, err := s.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Возвращаем найденный товар в формате JSON
	c.JSON(http.StatusOK, user)
}

func Create(c *gin.Context) {
	var body string
	var s user.Service

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := s.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func Update(c *gin.Context) {
	var body string
	var s user.Service

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := s.Update(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func Delete(c *gin.Context) {
	var s user.Service

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = s.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
