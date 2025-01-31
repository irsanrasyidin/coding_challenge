package handler

import (
	"coding_challenge/model"
	"coding_challenge/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RentHandler struct {
	rentUsecase usecase.RentUsecase
}

func (rentHandler *RentHandler) CreateRent(ctx *gin.Context) {
	var rent model.Orders
	if err := ctx.ShouldBindJSON(&rent); err != nil {
		fmt.Printf("rentHandler.ReadByIDRent(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}

	err := rentHandler.rentUsecase.CreateRent(&rent)
	if err != nil {
		fmt.Printf("rentHandler.CreateRent(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"Message": "Berhasil Membuat ID=" + strconv.Itoa(rent.OrderID),
	})
}

func (rentHandler *RentHandler) ReadAllRent(ctx *gin.Context) {
	rentData, err := rentHandler.rentUsecase.ReadAllRent()
	if err != nil {
		fmt.Printf("rentHandler.ReadAllRent(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rentData,
	})
}

func (rentHandler *RentHandler) ReadByIDRent(ctx *gin.Context) {
	id := ctx.Query("id")
	idInt, _ := strconv.Atoi(id)
	rentData, err := rentHandler.rentUsecase.ReadByIDRent(idInt)
	if err != nil {
		fmt.Printf("rentHandler.ReadByIDRent(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rentData,
	})
}

func (rentHandler *RentHandler) UpdateRent(ctx *gin.Context) {
	var rent *model.Orders
	if err := ctx.ShouldBindJSON(&rent); err != nil {
		fmt.Printf("rentHandler.ReadByIDRent(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}

	fmt.Println("AW",rent)

	err := rentHandler.rentUsecase.UpdateRent(rent)
	if err != nil {
		fmt.Printf("rentHandler.UpdateRent(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"Message": "Berhasil Update ID=" + strconv.Itoa(rent.OrderID),
	})
}

func (rentHandler *RentHandler) DeleteRent(ctx *gin.Context) {
	id := ctx.Query("id")
	idInt, _ := strconv.Atoi(id)
	err := rentHandler.rentUsecase.DeleteRent(idInt)
	if err != nil {
		fmt.Printf("rentHandler.ReadByIDRent(): %v", err.Error())
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

func NewRentHandler(srv *gin.Engine, rentUsecase usecase.RentUsecase) *RentHandler {
	rentHandler := &RentHandler{
		rentUsecase: rentUsecase,
	}

	rentGroup := srv.Group("/rent")
	{
		rentGroup.GET("/all", rentHandler.ReadAllRent)
		rentGroup.GET("/", rentHandler.ReadByIDRent)
		rentGroup.POST("/update", rentHandler.UpdateRent)
		rentGroup.POST("/create", rentHandler.CreateRent)
		rentGroup.DELETE("/", rentHandler.DeleteRent)
	}

	return rentHandler
}
