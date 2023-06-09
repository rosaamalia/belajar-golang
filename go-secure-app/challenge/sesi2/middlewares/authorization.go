package middlewares

import (
	"net/http"
	"sesi2/database"
	"sesi2/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		productId, err := strconv.Atoi(c.Param("productId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter.",
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := string(userData["role"].(string))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})

			return
		}

		if Product.UserID != userID {
			if userRole == "admin" {
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data.",
				})

			}

			return
		}

		c.Next()
	}
}

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		userData := c.MustGet("userData").(jwt.MapClaims)
		userRole := string(userData["role"].(string))

		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data.",
			})

			return
		}

		c.Next()
	}
}