package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"roastkuy-api/data/request"
	"roastkuy-api/data/response"
	"roastkuy-api/helper"
	"roastkuy-api/service"
)

var validate = validator.New()

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}

func (controller *AuthController) Login(ctx *gin.Context, loginRequest request.LoginRequest) {
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

func (controller *AuthController) HandleLogin(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.Login(ctx, loginRequest)
}

func (controller *AuthController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateAccountRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	err = validate.Struct(createUserRequest)
	if err != nil {
		// Handle validation errors with custom messages
		errors := helper.FormatValidationError(err)
		ctx.JSON(400, gin.H{
			"errors": errors,
		})
		return
	}

	controller.authService.Register(createUserRequest)
	// Get the login credentials from the registration request
	loginRequest := request.LoginRequest{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	// Login the user
	controller.Login(ctx, loginRequest)
}
