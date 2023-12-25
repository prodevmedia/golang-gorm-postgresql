package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supardi98/golang-gorm-postgres/app/controllers"
	"github.com/supardi98/golang-gorm-postgres/app/middleware"
	"gorm.io/gorm"
)

func AuthRoute(rg *gin.RouterGroup, dbConnection *gorm.DB) {
	authController := controllers.NewAuthController(dbConnection)

	router := rg.Group("/auth")

	router.POST("/register", authController.SignUpUser)
	router.POST("/login", authController.SignInUser)
	router.GET("/logout", middleware.DeserializeUser(), authController.LogoutUser)
	router.GET("/verifyemail/:verificationCode", authController.VerifyEmail)
	router.POST("/forgotpassword", authController.ForgotPassword)
	router.PATCH("/resetpassword/:resetToken", authController.ResetPassword)
}
