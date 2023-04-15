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

func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Debug().Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}
	
	commentId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Debug().Where("id = ?", commentId).Find(&Comment).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	// assign id dari token jwt ke data inputan
	Comment.UserID = userId

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}
	
	commentId, _ := strconv.Atoi(c.Param("id"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId
	Comment.ID = uint(commentId)

	var photoId uint
	model := models.Comment{}
	db.Where("id = ?", commentId).Find(&model).Select("photo_id").Scan(&photoId)
	Comment.PhotoID = photoId

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}
	
	photoId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("id = ?", photoId).Delete(&Comment).Error

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