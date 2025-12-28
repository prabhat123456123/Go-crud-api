package main

import (
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/routes"
"go-crud-api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// DB Connection
	config.ConnectDB()

	// Auto migrate
	config.DB.AutoMigrate(&models.User{})

	// Routes
	routes.UserRoutes(r)
	routes.AuthRoutes(r)
r.Use(middlewares.ErrorHandler())

	r.Run(":8080")
}
