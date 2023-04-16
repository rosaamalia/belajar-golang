package router

import (
	"sesi4/controllers"
	"sesi4/middlewares"

	_ "sesi4/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
// @description This is an api for MyGram Project
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/auth")
	{
		// Register user
		userRouter.POST("/register", controllers.UserRegister)
		// Login user
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		// Get All Photos
		photoRouter.GET("/", controllers.GetAllPhotos)
		// Create Photo
		photoRouter.POST("/", controllers.CreatePhoto)
		// Get Photo By ID
		photoRouter.GET("/:id", controllers.GetPhotoById)
		// Update Photo
		photoRouter.PUT("/:id", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		// Delete Photo
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		// Get All Comments
		commentRouter.GET("/", controllers.GetAllComments)
		// Create Comment
		commentRouter.POST("/", controllers.CreateComment)
		// Get Comment by ID
		commentRouter.GET("/:id", controllers.GetCommentById)
		// Update Comment
		commentRouter.PUT("/:id", middlewares.CommentAuthorization(), controllers.UpdateComment)
		// Delete Comment
		commentRouter.DELETE("/:id", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social_medias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		// Get All Social Media
		socialMediaRouter.GET("/", controllers.GetAllSocialMedias)
		// Create Social Media
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		// Get Social Media by ID
		socialMediaRouter.GET("/:id", controllers.GetSocialMediaById)
		// Update Social Media
		socialMediaRouter.PUT("/:id", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		// Delete Social Media
		socialMediaRouter.DELETE("/:id", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}