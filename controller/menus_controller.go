package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roastkuy-api/data/response"
	"roastkuy-api/service"
	"strconv"
)

type MenusController struct {
	menusService service.MenusService
}

func NewMenusController(menusService service.MenusService) *MenusController {
	return &MenusController{menusService: menusService}
}

func (controller *MenusController) FindByOutletId(ctx *gin.Context) {
	outletId := ctx.Param("outletId")
	id, err := strconv.Atoi(outletId)
	if err != nil {
		webResponse := response.Response{
			Code:    "bad_request",
			Message: "Bad Request",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	outletResponse, err := controller.menusService.FindByOutletId(id)
	if err != nil || outletResponse == nil {
		ctx.JSON(http.StatusNotFound, response.DataResponse{
			Message: "Data tidak ditemukan!",
			Data:    nil,
		})
		return
	}

	webResponse := response.DataResponse{
		Message: "Data ditemukan!",
		Data:    outletResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MenusController) FindById(ctx *gin.Context) {
	menuId := ctx.Param("menuId")
	id, err := strconv.Atoi(menuId)
	if err != nil {
		webResponse := response.Response{
			Code:    "bad_request",
			Message: "Bad Request",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	outletResponse, err := controller.menusService.FindById(id)
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
