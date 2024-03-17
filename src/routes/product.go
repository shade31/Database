package routes

import (
	"Database/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductsGroupRouter(baseRouter *gin.RouterGroup) {
	products := baseRouter.Group("/products")

	products.GET("/get", controllers.GetAllProducts)
	products.GET("/get/:id", controllers.GetProductById)
	products.POST("/create", controllers.CreateProduct)
	products.PATCH("/update/:id", controllers.UpdateProduct)
	products.DELETE("/delete/:id", controllers.DeleteProduct)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	ProductsGroupRouter(versionRouter)

	return r
}
