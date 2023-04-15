package middlewares

import (
	"net/http"
	"sesi4/database"
	"sesi4/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"error": "Bad Request",
				"message": "Invalid parameter",
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		model := models.Photo{}
		
		/**
		 * memastikan data dengan id (sesuai parameter)
		 * memiliki user_id yg sama dengan userId
		 */
		err = db.Debug().Where("id = ?", id).Find(&model).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"error": "Data not found",
				"message": "Data doesn't exist",
			})

			return
		}

		if model.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to modify this data",
			})

			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"error": "Bad Request",
				"message": "Invalid parameter",
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		model := models.Comment{}
		
		err = db.Debug().Where("id = ?", id).Find(&model).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"error": "Data not found",
				"message": "Data doesn't exist",
			})

			return
		}

		if model.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to modify this data",
			})

			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"error": "Bad Request",
				"message": "Invalid parameter",
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		model := models.SocialMedia{}

		err = db.Debug().Where("id = ?", id).Find(&model).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"error": "Data not found",
				"message": "Data doesn't exist",
			})

			return
		}

		if model.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to modify this data",
			})

			return
		}

		c.Next()
	}
}