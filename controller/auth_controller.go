package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roastkuy-api/data/request"
	"roastkuy-api/data/response"
	"roastkuy-api/helper"
	"roastkuy-api/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, expiresIn, err_token := controller.authService.Login(loginRequest)

	if err_token != nil {
		switch err_token.Error() {
		case "ACCOUNT_NOT_FOUND":
			webResponse := response.Response{
				Code:    "unauthorized",
				Message: "Maaf, anda tidak terdaftar dalam sistem kami",
				Data:    nil,
			}
			ctx.JSON(http.StatusNotFound, webResponse)

		case "WRONG_PASSWORD":
			webResponse := response.Response{
				Code:    "unauthorized",
				Message: "Kata sandi yang anda masukan salah",
				Data:    nil,
			}
			ctx.JSON(http.StatusUnauthorized, webResponse)

		default:
			webResponse := response.Response{
				Code:    "error",
				Message: "An error occurred",
				Data:    nil,
			}
			ctx.JSON(http.StatusInternalServerError, webResponse)
		}
	} else {
		expiresInStr := expiresIn.Format("2006-01-02 15:04:05")

		resp := response.LoginResponse{
			ExpiresIn: expiresInStr,
			Token:     token,
			IsActive:  loginRequest.Email,
		}

		webResponse := response.Response{
			Code:    "success",
			Message: "Login berhasil!",
			Data:    resp,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}

}

func (controller *AuthController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateAccountRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.authService.Register(createUserRequest)
	webResponse := response.Response{
		Code:    "Sukses",
		Message: "Success Create User",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
