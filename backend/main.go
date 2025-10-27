package main

import (
	"go_crud_with_gin_framework/config"
	"go_crud_with_gin_framework/models"
	"go_crud_with_gin_framework/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Product{}, &models.User{})
	r := gin.Default()

	routes.ProductRoutes(r)
	routes.UserRoutes(r)

	r.Run(":8080")
}
