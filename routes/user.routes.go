package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supardi98/golang-gorm-postgres/app/controllers"
	"github.com/supardi98/golang-gorm-postgres/app/middleware"
	"gorm.io/gorm"
)

func UserRoute(rg *gin.RouterGroup, dbConnection *gorm.DB) {
	userController := controllers.NewUserController(dbConnection)

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), userController.GetMe)
}
