package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prodevmedia/golang-gorm-postgres/app/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	ResponseWithSuccess(ctx, http.StatusOK, gin.H{"user": currentUser}, "Profile retrieved successfully")
}
