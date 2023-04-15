package controllers

import (
	"net/http"
	"sesi4/database"
	"sesi4/helpers"
	"sesi4/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Debug().Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}
	
	socialMediaId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Debug().Where("id = ?", socialMediaId).Find(&SocialMedia).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	// assign id dari token jwt ke data inputan
	SocialMedia.UserID = userId

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}
	
	socialMediaId, _ := strconv.Atoi(c.Param("id"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userId
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}
	
	photoId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Where("id = ?", photoId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H {
		"message": "Data is deleted successfully",
	})
}