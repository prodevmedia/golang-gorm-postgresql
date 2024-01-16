package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prodevmedia/golang-gorm-postgres/app/models"
	"github.com/prodevmedia/golang-gorm-postgres/app/utils"
	"gorm.io/gorm"
)

type ProfileController struct {
	DB *gorm.DB
}

func NewProfileController(DB *gorm.DB) ProfileController {
	return ProfileController{DB}
}

func (pc *ProfileController) GetProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	ResponseWithSuccess(ctx, http.StatusOK, gin.H{"user": currentUser.Response()})
}

func (pc *ProfileController) UpdateProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateProfileInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userToUpdate := models.User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	pc.DB.Model(&currentUser).Updates(userToUpdate)

	ResponseWithSuccess(ctx, http.StatusOK, currentUser.Response())
}

func (pc *ProfileController) UpdatePassword(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePasswordInput
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

	userToUpdate := models.User{
		Password: hashedPassword,
	}

	pc.DB.Model(&currentUser).Updates(userToUpdate)

	ResponseWithSuccess(ctx, http.StatusOK, "Password updated successfully")
}

func (pc *ProfileController) UploadAvatar(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	file, err := ctx.FormFile("file")
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	avatarPath, err := utils.UploadFile(file, "avatar")
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userToUpdate := models.User{
		Avatar: avatarPath,
	}

	pc.DB.Model(&currentUser).Updates(userToUpdate)

	ResponseWithSuccess(ctx, http.StatusOK, currentUser.Response())
}
