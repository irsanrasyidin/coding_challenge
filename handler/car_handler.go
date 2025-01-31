package handler

import (
	"coding_challenge/model"
	"coding_challenge/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarsHandler struct {
	carsUsecase usecase.CarsUsecase
}

func (carsHandler *CarsHandler) CreateCars(ctx *gin.Context) {
	var cars model.Cars
	if err := ctx.ShouldBindJSON(&cars); err != nil {
		fmt.Printf("carsHandler.ReadByIDRent(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}

	fmt.Println("handlers", cars)

	err := carsHandler.carsUsecase.CreateCars(&cars)
	if err != nil {
		fmt.Printf("carsHandler.CreateCars(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"Message": "Berhasil Membuat ID=" + strconv.Itoa(cars.CarID),
	})
}

func (carsHandler *CarsHandler) ReadAllCars(ctx *gin.Context) {
	carsData, err := carsHandler.carsUsecase.ReadAllCars()
	if err != nil {
		fmt.Printf("carsHandler.ReadAllCars(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    carsData,
	})
}

func (carsHandler *CarsHandler) ReadByIDCars(ctx *gin.Context) {
	id := ctx.Query("id")
	fmt.Println(id)
	idInt, _ := strconv.Atoi(id)
	carsData, err := carsHandler.carsUsecase.ReadByIDCars(idInt)
	if err != nil {
		fmt.Printf("carsHandler.ReadByIDCars(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    carsData,
	})
}

func (carsHandler *CarsHandler) UpdateCars(ctx *gin.Context) {
	var cars model.Cars
	if err := ctx.ShouldBindJSON(&cars); err != nil {
		fmt.Printf("carsHandler.ReadByIDCars(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"Message": "Berhasil Update ID=" + strconv.Itoa(cars.CarID),
	})
}

func (carsHandler *CarsHandler) DeleteCars(ctx *gin.Context) {
	id := ctx.Query("id")
	idInt, _ := strconv.Atoi(id)
	err := carsHandler.carsUsecase.DeleteCars(idInt)
	if err != nil {
		fmt.Printf("carsHandler.ReadByIDCars(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"Message": "Berhasil Delete ID=" + id,
	})
}

func NewCarsHandler(srv *gin.Engine, carsUsecase usecase.CarsUsecase) *CarsHandler {
	carsHandler := &CarsHandler{
		carsUsecase: carsUsecase,
	}

	carsGroup := srv.Group("/cars")
	{
		carsGroup.GET("/all", carsHandler.ReadAllCars)
		carsGroup.GET("/", carsHandler.ReadByIDCars)
		carsGroup.POST("/update", carsHandler.UpdateCars)
		carsGroup.POST("/create", carsHandler.CreateCars)
		carsGroup.DELETE("/", carsHandler.DeleteCars)
	}
	return carsHandler
}
