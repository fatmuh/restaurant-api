package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roastkuy-api/controller"
)

func NewRouter(authController *controller.AuthController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome home"})
	})

	baseRouter := router.Group("/api")

	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.HandleLogin)

	return router
}
