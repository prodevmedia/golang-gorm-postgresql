package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supardi98/golang-gorm-postgres/app/controllers"
	"github.com/supardi98/golang-gorm-postgres/app/middleware"
	"gorm.io/gorm"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(dbConnection *gorm.DB) UserRouteController {
	userController := controllers.NewUserController(dbConnection)

	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
