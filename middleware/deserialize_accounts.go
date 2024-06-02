package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	config2 "roastkuy-api/config"
	"roastkuy-api/helper"
	"roastkuy-api/repository"
	"roastkuy-api/utils"
	"strconv"
	"strings"
)

func DeserializeAccounts(accountsRepository repository.AccountsRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Failed", "message": "Unauthorized"})
			return
		}

		config, _ := config2.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Failed", "message": err.Error()})
		}

		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helper.ErrorPanic(err_id)
		result, err := accountsRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Failed", "message": "The user belonging to this token no longer"})
			return
		}

		ctx.Set("currentUser", result)
		ctx.Next()
	}
}
