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

func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Debug().Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func GetPhotoById(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}
	
	photoId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Debug().Where("id = ?", photoId).Find(&Photo).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	// assign id dari token jwt ke data inputan
	Photo.UserID = userId

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}
	
	photoId, _ := strconv.Atoi(c.Param("id"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}
	
	photoId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

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