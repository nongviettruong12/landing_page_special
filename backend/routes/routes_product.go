package routes

import (
	"go_crud_with_gin_framework/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProduct)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
}
