package router

import (
	"sesi2/controllers"
	"sesi2/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProductbyID)

		productRouter.PUT("/:productId", middlewares.AdminAuthorization(), controllers.UpdateProduct)

		productRouter.DELETE("/:productId", middlewares.AdminAuthorization(), controllers.DeleteProduct)
	}

	return r
}