package controllers

import (
	"go_crud_with_gin_framework/config"
	"go_crud_with_gin_framework/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /products
func GetUser(c *gin.Context) {
	var products []models.User
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// GET /products/:id
func GetUserById(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy người dùng"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// POST /products
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// PUT /products/:id
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy người dùng"})
		return
	}

	var updated models.User
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&user).Updates(updated)
	c.JSON(http.StatusOK, user)
}

// DELETE /products/:id
func DeleteUser(c *gin.Context) {
	config.DB.Delete(&models.User{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Đã xoá người dùng"})
}
