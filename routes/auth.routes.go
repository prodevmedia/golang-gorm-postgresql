package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supardi98/golang-gorm-postgres/app/controllers"
	"github.com/supardi98/golang-gorm-postgres/app/middleware"
	"gorm.io/gorm"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(dbConnection *gorm.DB) AuthRouteController {
	authController := controllers.NewAuthController(dbConnection)

	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
	router.GET("/verifyemail/:verificationCode", rc.authController.VerifyEmail)
	router.POST("/forgotpassword", rc.authController.ForgotPassword)
	router.PATCH("/resetpassword/:resetToken", rc.authController.ResetPassword)
}
