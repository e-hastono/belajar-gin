package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"belajar-gin-cars/models"

	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @Summary Post car info
// @Description Post car info
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 201 {object} models.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	var newCar models.Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := newCar.CreateCar()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

// UpdateCar godoc
// @Summary Update car info of given id
// @Description Update car information for a given id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Id of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [put]
func UpdateCar(ctx *gin.Context) {
	rawCarID := ctx.Param("carID")
	carID, err := strconv.Atoi(rawCarID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var updatedCar = models.Car{ID: carID}

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = updatedCar.UpdateCar()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
	})
}

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all cars
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(ctx *gin.Context) {
	cars, err := models.GetAllCars()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cars": cars,
	})
}

// GetCar godoc
// @Summary Get car info for a given id
// @Description Get information of a car with a given id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetCar(ctx *gin.Context) {
	rawCarID := ctx.Param("carID")
	carID, err := strconv.Atoi(rawCarID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	car, err := models.GetCar(carID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}

// DeleteCar godoc
// @Summary Delete car with given id
// @Description Delete car information based on given id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Id of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{id} [delete]
func DeleteCar(ctx *gin.Context) {
	rawCarID := ctx.Param("carID")
	carID, err := strconv.Atoi(rawCarID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	carID, err = models.DeleteCar(carID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
	})
}
