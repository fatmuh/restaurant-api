package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roastkuy-api/data/response"
	"roastkuy-api/middleware"
	"roastkuy-api/service"
)

type PromosController struct {
	promosService service.PromosService
}

func NewPromosController(promosService service.PromosService) *PromosController {
	return &PromosController{promosService: promosService}
}

func (controller *PromosController) FindAll(ctx *gin.Context) {
	userID, exists := middleware.GetUserID(ctx)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}

	outletResponse := controller.promosService.FindAll(userID)
	webResponse := response.DataResponse{
		Message: "Data ditemukan!",
		Data:    outletResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
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
