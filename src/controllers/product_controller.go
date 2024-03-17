package controllers

import (
	"Database/src/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	productsInterface, err := models.SelectQuery(models.Database)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	products := productsInterface
	c.JSON(http.StatusOK, gin.H{"message": "Startups fetched successfully", "status": "success", "data": products})
}

func GetProductById(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}
	productInfo, err := models.SelectQueryById(models.Database, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Startup fetched successfully", "status": "success", "data": productInfo})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		log.Println("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	id, err := strconv.Atoi(product.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	if err := models.InsertQuery(models.Database, models.Product{Id: id, Name: product.Name, Price: product.Price, Count: product.Count}); err != nil {
		log.Println("Failed to create product:", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	// log.Println("Product created successfully - ID:", product.Id)
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully"})
}

// func UpdateStartup(c *gin.Context) {
// 	startupID := c.Param("id")

// 	var updatedStartup *models.Startup
// 	if err := c.ShouldBindJSON(&updatedStartup); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
// 		return
// 	}

// 	updatedStartup, err := updatedStartup.UpdateStartup(startupID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Startup updated successfully", "data": updatedStartup})
// }

// func DeleteStartup(c *gin.Context) {
// 	startupID := c.Param("id")
// 	if startupID == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
// 		return
// 	}

// 	err := models.DeleteStartup(startupID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Startup deleted successfully", "status": "success", "data": nil})
// }
