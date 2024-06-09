package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roastkuy-api/controller"
	"roastkuy-api/middleware"
	"roastkuy-api/repository"
)

func NewRouter(accountRepository repository.AccountsRepository, authController *controller.AuthController, outletsController *controller.OutletsController, menusController *controller.MenusController, promosController *controller.PromosController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome home"})
	})

	baseRouter := router.Group("/api")

	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.HandleLogin)

	outletRouter := baseRouter.Group("/outlet")
	outletRouter.GET("", outletsController.FindAll)
	outletRouter.GET("/:slug", outletsController.FindBySlug)

	menuRouter := baseRouter.Group("/menu")
	menuRouter.GET("/:outletId", menusController.FindByOutletId)

	menuDetailRouter := baseRouter.Group("/menu-detail")
	menuDetailRouter.GET("/:menuId", menusController.FindById)

	promoRouter := baseRouter.Group("/promo")
	promoRouter.GET("", middleware.DeserializeAccounts(accountRepository), promosController.FindAll)

	promoRegularRouter := baseRouter.Group("/promo-regular")
	promoRegularRouter.GET("", promosController.FindRegular)
	return router
}
