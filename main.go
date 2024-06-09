package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
	"roastkuy-api/config"
	"roastkuy-api/controller"
	"roastkuy-api/helper"
	"roastkuy-api/repository"
	"roastkuy-api/router"
	"roastkuy-api/service"
)

func main() {
	log.Info().Msg("Starting server...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	// Database
	db := config.DatabaseConnection(&loadConfig)
	validate := validator.New()

	// Repository
	accountsRepository := repository.NewAccountsRepositoryImpl(db)
	outletsRepository := repository.NewOutletsRepositoryImpl(db)

	// Service
	authenticationService := service.NewAuthenticationServiceImpl(accountsRepository, validate, db)
	outletsService := service.NewOutletsServiceImpl(outletsRepository, validate)

	// Controller
	authenticationController := controller.NewAuthController(authenticationService)
	outletsController := controller.NewOutletsController(outletsService)

	// Router
	routes := router.NewRouter(authenticationController, outletsController)

	routes.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
