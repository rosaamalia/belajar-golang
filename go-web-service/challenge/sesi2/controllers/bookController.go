package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID 	 string `json:id`
	Title 	 string `json:title`
	Author 	 string	`json:author`
	Desc	 string `json:desc`
}

var BookData = []Book{}

/**
 * Mendapatkan semua data
 */
func GetAllBooks (ctx *gin.Context) {
	if len(BookData) > 0 {
		ctx.JSON(http.StatusOK, gin.H {
			"data": BookData,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H {
			"data": "No book saved",
		})	
	}
}

/**
 * Mendapatkan data berdasarkan ID
 */
func GetBookByID (ctx *gin.Context) {
	BookID := ctx.Param("BookID")
	var book Book

	for i, item := range BookData {
		if BookID == item.BookID {
			book = BookData[i]
			break
		}
	}

	ctx.JSON(http.StatusOK, gin.H {
		"data": book,
	})
}

/**
 * Menambahkan data baru
 */
func CreateBook(ctx *gin.Context) {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// menambahkan data
	book.BookID = fmt.Sprintf("%d", len(BookData)+1)
	BookData = append(BookData, book)
	
	// mengembalikan response
	ctx.JSON(http.StatusCreated, gin.H{
		"data": "Created",
	})
}

/**
 * Mengubah data
 */
func UpdateBook(ctx *gin.Context) {
	BookID := ctx.Param("BookID")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookData {
		if BookID == book.BookID {
			condition = true
			BookData[i] = updatedBook
			BookData[i].BookID = BookID
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"data": "Data not found.",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"data": "Updated",
	})
}

/**
 * Menghapus data
 */
func DeleteBook(ctx *gin.Context) {
	BookID := ctx.Param("BookID")
	condition := false

	var bookIndex int

	for i, book := range BookData {
		if BookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data": "Data not found",
		})

		return
	}

	copy(BookData[bookIndex:], BookData[bookIndex+1:])
	BookData[len(BookData)-1] = Book{}
	BookData = BookData[:len(BookData)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Deleted",
	})
}