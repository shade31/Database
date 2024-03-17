package controllers

import (
	"Database/src/models"
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
	c.JSON(http.StatusOK, gin.H{"message": "Products fetched successfully", "status": "success", "data": products})
}

func GetProductById(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product ID is required"})
		return
	}
	productInfo, err := models.SelectQueryById(models.Database, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product fetched successfully", "status": "success", "data": productInfo})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	if err := models.InsertQuery(models.Database, product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Product saved successfully"})
}

func UpdateProduct(c *gin.Context) {
	var updateData struct {
		Price int `json:"price"`
		Count int `json:"count"`
	}

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid JSON data"})
		return
	}
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Product ID is required"})
		return
	}

	if err := models.UpdateQuery(models.Database, id, updateData.Price, updateData.Count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product fetched successfully", "status": "success"})
}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Product ID is required"})
		return
	}

	if err := models.DeleteQuery(models.Database, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Startup deleted successfully"})
}
