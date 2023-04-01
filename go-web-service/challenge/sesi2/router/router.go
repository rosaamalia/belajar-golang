package router

import (
	"sesi2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	/**
	 * GET
	 */
	router.GET("/book", controllers.GetAllBooks)
	/**
	 * POST
	 */
	router.POST("/book", controllers.CreateBook)
	/**
	 * GET BY ID
	 */
	router.GET("/book/:BookID", controllers.GetBookByID)
	/**
	 * PUT
	 */
	router.PUT("/book/:BookID", controllers.UpdateBook)
	/**
	 * DELETE
	 */
	router.DELETE("/book/:BookID", controllers.DeleteBook)

	return router
}