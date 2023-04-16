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

// GetAllSocialMedias godoc
// @Summary Get all details
// @Description Get all social_medias data
// @Tags social_medias
// @Accept json
// @Produce json
// @Security bearerAuth
// @Success 200 {object} models.SocialMedia
// @Router /social_medias [get]
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

// GetSocialMediaById godoc
// @Summary Get detail of a data
// @Description Get detail of a social_medias data
// @Tags social_medias
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param Id path int true "ID of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /social_medias/{id} [get]
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

// CreateSocialMedia godoc
// @Summary Create a data
// @Description Create new social_media data
// @Tags social_medias
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param requestBody body models.SocialMedia true "create social media"
// @Success 200 {object} models.SocialMedia
// @Router /social_medias/ [post]
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

// UpdateSocialMedia godoc
// @Summary Update a data
// @Description Update a social_media data
// @Tags social_medias
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param Id path int true "ID of the social media"
// @Param requestBody body models.SocialMedia true "update social media"
// @Success 200 {object} models.SocialMedia
// @Router /social_medias/{id} [put]
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

// DeleteSocialMedia godoc
// @Summary Delete a data
// @Description Delete a social_media data
// @Tags social_medias
// @Accept json
// @Produce json
// @Security bearerAuth

// @Param Id path int true "ID of the social_media"
// @Success 200 {object} models.SocialMedia
// @Router /social_medias/{id} [delete]
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