package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int 	 `json:"price"`
}

var carData = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car
	
	/**
	 * Request body di-binding ke sebuah tipe menggunakan
	 * model binding. Gin dapat melakukan binding dari
	 * JSON, XML, YAML, dan standard form values
	 * https://gin-gonic.com/docs/examples/binding-and-validation/
	 */

	// Validasi apakah data JSON pada body tidak ada error
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// menambahkan data
	newCar.CarID = fmt.Sprintf("%d", len(carData)+1)
	carData = append(carData, newCar)

	/**
	 * Untuk mengembalikan response dalam bentuk JSON,
	 * fungsi ctx.JSON membutuhkan dua parameter yaitu
	 * status code dan response
	 */
	
	// mengembalikan response
	ctx.JSON(http.StatusCreated, gin.H{
		"data": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	// mengambil parameter
	carID := ctx.Param("carID")
	condition := false
	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// mencari data car yang sama dengan carID
	for i, car := range carData {
		if carID == car.CarID {
			condition = true
			// data lama diupdate sesuai dengan data baru
			carData[i] = updatedCar
			carData[i].CarID = carID
		}
	}

	// jika data yang dicari tidak ditemukan
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"data": "Data not found.",
		})

		return
	}

	// mengembalikan response
	ctx.JSON(http.StatusOK, gin.H {
		"data": "Data successfully updated.",
	})
}