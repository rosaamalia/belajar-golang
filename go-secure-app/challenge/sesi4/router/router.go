package router

import (
	"sesi4/controllers"
	"sesi4/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/auth")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhotos)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/:id", controllers.GetPhotoById)
		photoRouter.PUT("/:id", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComments)
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/:id", controllers.GetCommentById)
		commentRouter.PUT("/:id", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:id", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social_medias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", controllers.GetAllSocialMedias)
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/:id", controllers.GetSocialMediaById)
		socialMediaRouter.PUT("/:id", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}