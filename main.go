package main

import (
	"Database/src/models"
	"Database/src/routes"
	"Database/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	utils.LoadEnv()
	models.OpenDatabaseConnection()
	defer models.CloseDatabaseConnection()
	// models.SelectQuery(models.Database)

	router := gin.Default()

	versionRouter := router.Group("/api/v1")
	routes.ProductsGroupRouter(versionRouter)

	// Simple "Hello Gin" route
	router.GET("/gin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin!"})
	})

	// serve and listen to localhost:8080
	router.Run(":8080")
}
