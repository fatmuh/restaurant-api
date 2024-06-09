package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roastkuy-api/data/response"
	"roastkuy-api/service"
)

type OutletsController struct {
	outletsService service.OutletsService
}

func NewOutletsController(outletsService service.OutletsService) *OutletsController {
	return &OutletsController{outletsService: outletsService}
}

func (controller *OutletsController) FindAll(ctx *gin.Context) {
	outletResponse := controller.outletsService.FindAll()
	webResponse := response.DataResponse{
		Message: "Data ditemukan!",
		Data:    outletResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OutletsController) FindBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	outletResponse, err := controller.outletsService.FindBySlug(slug)
	if err != nil {
		webResponse := response.DataResponse{
			Message: "Data tidak ditemukan!",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.DataResponse{
		Message: "Data ditemukan!",
		Data:    outletResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
