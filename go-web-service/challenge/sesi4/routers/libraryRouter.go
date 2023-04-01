package routers

import (
	"sesi4/controllers"
	"sesi4/database"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	db := database.StartDB()
	controllers := controllers.New(db)

	router.GET("/books", controllers.GetAllBooks)

	router.GET("/books/:bookID", controllers.GetBookbyID)

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:bookID", controllers.UpdateBookbyID)

	router.DELETE("/books/:bookID", controllers.DeleteBookbyID)

	return router
}