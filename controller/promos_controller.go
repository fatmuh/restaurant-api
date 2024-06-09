package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roastkuy-api/data/response"
	"roastkuy-api/service"
)

type PromosController struct {
	promosService service.PromosService
}

func NewPromosController(promosService service.PromosService) *PromosController {
	return &PromosController{promosService: promosService}
}

func (controller *PromosController) FindRegular(ctx *gin.Context) {
	outletResponse := controller.promosService.FindRegular()
	webResponse := response.DataResponse{
		Message: "Data ditemukan!",
		Data:    outletResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
