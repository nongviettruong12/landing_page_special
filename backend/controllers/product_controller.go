package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go_crud_with_gin_framework/config"
    "go_crud_with_gin_framework/models"
)

// GET /products
func GetProducts(c *gin.Context) {
    var products []models.Product
    config.DB.Find(&products)
    c.JSON(http.StatusOK, products)
}

// GET /products/:id
func GetProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
        return
    }
    c.JSON(http.StatusOK, product)
}

// POST /products
func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&product)
    c.JSON(http.StatusCreated, product)
}

// PUT /products/:id
func UpdateProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
        return
    }

    var updated models.Product
    if err := c.ShouldBindJSON(&updated); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Model(&product).Updates(updated)
    c.JSON(http.StatusOK, product)
}

// DELETE /products/:id
func DeleteProduct(c *gin.Context) {
    config.DB.Delete(&models.Product{}, c.Param("id"))
    c.JSON(http.StatusOK, gin.H{"message": "Đã xoá sản phẩm"})
}
