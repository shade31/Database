package routes

import (
	"Database/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductsGroupRouter(baseRouter *gin.RouterGroup) {
	products := baseRouter.Group("/products")

	products.GET("/all", controllers.GetAllProducts)
	products.GET("/get/:id", controllers.GetProductById)
	products.POST("/create", controllers.CreateProduct)
	// products.PATCH("/update", controllers.UpdateStartup)
	// products.DELETE("/delete/:id", controllers.DeleteStartup)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	ProductsGroupRouter(versionRouter)

	return r
}
