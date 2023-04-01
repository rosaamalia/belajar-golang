package routers

import (
	"3-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	/**
	 * POST
	 */
	router.POST("/cars", controllers.CreateCar)
	/**
	 * PUT
	 */
	router.PUT("/cars/:carID", controllers.UpdateCar)

	return router
}

