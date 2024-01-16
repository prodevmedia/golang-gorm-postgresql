package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prodevmedia/golang-gorm-postgres/app/controllers"
	"github.com/prodevmedia/golang-gorm-postgres/app/middleware"
	"gorm.io/gorm"
)

func ProfileRoute(rg *gin.RouterGroup, dbConnection *gorm.DB) {
	profileController := controllers.NewProfileController(dbConnection)

	router := rg.Group("profile")
	router.GET("/", middleware.DeserializeUser(), profileController.GetProfile)
	router.PUT("/", middleware.DeserializeUser(), profileController.UpdateProfile)
	router.PUT("/password", middleware.DeserializeUser(), profileController.UpdatePassword)
	router.PUT("/avatar", middleware.DeserializeUser(), profileController.UploadAvatar)
}
