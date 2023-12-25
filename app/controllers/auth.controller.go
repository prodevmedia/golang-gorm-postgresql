package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/supardi98/golang-gorm-postgres/app/models"
	"github.com/supardi98/golang-gorm-postgres/app/utils"
	"github.com/supardi98/golang-gorm-postgres/config"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// [...] SignUp User
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ResponseWithError(ctx, http.StatusBadRequest, "Passwords do not match")
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ResponseWithError(ctx, http.StatusBadGateway, err.Error())
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  false,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ResponseWithError(ctx, http.StatusConflict, "User with that email already exists")
		return
	} else if result.Error != nil {
		ResponseWithError(ctx, http.StatusBadGateway, "Something bad happened")
		return
	}

	config, _ := config.LoadConfig(".")

	// Generate Verification Code
	code := randstr.String(20)

	verification_code := utils.Encode(code)

	// Update User in Database
	newUser.VerificationCode = verification_code
	ac.DB.Save(newUser)

	var firstName = newUser.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// ? Send Email
	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	utils.SendEmail(&newUser, &emailData, "verificationCode.html")

	message := "We sent an email with a verification code to " + newUser.Email

	ResponseWithSuccess(ctx, http.StatusCreated, message)
}

// [...] Verify Email
func (ac *AuthController) VerifyEmail(ctx *gin.Context) {

	code := ctx.Params.ByName("verificationCode")
	verification_code := utils.Encode(code)

	var updatedUser models.User
	result := ac.DB.First(&updatedUser, "verification_code = ?", verification_code)
	if result.Error != nil {
		ResponseWithError(ctx, http.StatusBadRequest, "Invalid verification code or user doesn't exists")
		return
	}

	if updatedUser.Verified {
		ResponseWithError(ctx, http.StatusConflict, "User already verified")
		return
	}

	updatedUser.VerificationCode = ""
	updatedUser.Verified = true
	ac.DB.Save(&updatedUser)

	ResponseWithSuccess(ctx, http.StatusOK, "Email verified successfully")
}

// [...] SignIn User
func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ResponseWithError(ctx, http.StatusBadRequest, "Invalid email or Password")
		return
	}

	if !user.Verified {
		ResponseWithError(ctx, http.StatusForbidden, "Please verify your email")
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, "Invalid email or Password")
		return
	}

	config, _ := config.LoadConfig(".")

	// Generate Token
	token, err := utils.GenerateToken(config.TokenExpiresIn, user.ID, config.TokenSecret)
	if err != nil {
		ResponseWithError(ctx, http.StatusBadGateway, err.Error())
		return
	}

	ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)

	ResponseWithSuccess(ctx, http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

// [...] SignOut User
func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ResponseWithSuccess(ctx, http.StatusOK, "User logged out successfully")
}

func (ac *AuthController) ForgotPassword(ctx *gin.Context) {
	var payload *models.ForgotPasswordInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := "You will receive a reset email if user with that email exist"

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ResponseWithError(ctx, http.StatusBadRequest, "Invalid email or Password")
		return
	}

	if !user.Verified {
		ResponseWithError(ctx, http.StatusUnauthorized, "Account not verified")
		return
	}

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// Generate Verification Code
	resetToken := randstr.String(20)

	passwordResetToken := utils.Encode(resetToken)
	user.PasswordResetToken = passwordResetToken
	user.PasswordResetAt = time.Now().Add(time.Minute * 15)
	ac.DB.Save(&user)

	var firstName = user.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// ? Send Email
	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "/resetpassword/" + resetToken,
		FirstName: firstName,
		Subject:   "Your password reset token (valid for 10min)",
	}

	utils.SendEmail(&user, &emailData, "resetPassword.html")

	ResponseWithSuccess(ctx, http.StatusOK, message)
}

func (ac *AuthController) ResetPassword(ctx *gin.Context) {
	var payload *models.ResetPasswordInput
	resetToken := ctx.Params.ByName("resetToken")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseWithSuccess(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ResponseWithSuccess(ctx, http.StatusBadRequest, "Passwords do not match")
		return
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	passwordResetToken := utils.Encode(resetToken)

	var updatedUser models.User
	result := ac.DB.First(&updatedUser, "password_reset_token = ? AND password_reset_at > ?", passwordResetToken, time.Now())
	if result.Error != nil {
		ResponseWithError(ctx, http.StatusBadRequest, "The reset token is invalid or has expired")
		return
	}

	updatedUser.Password = hashedPassword
	updatedUser.PasswordResetToken = ""
	ac.DB.Save(&updatedUser)

	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

	ResponseWithSuccess(ctx, http.StatusOK, "Password data updated successfully")
}
