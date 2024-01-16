package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prodevmedia/golang-gorm-postgres/app/controllers"
	"github.com/prodevmedia/golang-gorm-postgres/app/models"
	"github.com/prodevmedia/golang-gorm-postgres/app/utils"
	"github.com/prodevmedia/golang-gorm-postgres/config"
	"github.com/prodevmedia/golang-gorm-postgres/database"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		// cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}
		//  else if err == nil {
		// 	token = cookie
		// }

		if token == "" {
			controllers.ResponseWithError(ctx, http.StatusUnauthorized, "You are not logged in")
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			controllers.ResponseWithError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		var user models.User
		result := database.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			controllers.ResponseWithError(ctx, http.StatusUnauthorized, "You are not logged in")
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
