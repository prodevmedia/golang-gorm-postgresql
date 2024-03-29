package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prodevmedia/golang-gorm-postgres/app/models"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreatePostRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newPost := models.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Image:   payload.Image,
		UserID:  currentUser.ID,
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ResponseWithError(ctx, http.StatusConflict, "Post with that title already exists")
			return
		}
		ResponseWithError(ctx, http.StatusBadGateway, result.Error.Error())
		return
	}

	ResponseWithSuccess(ctx, http.StatusCreated, newPost, "Post created successfully")
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	postId := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var updatedPost models.Post
	result := pc.DB.First(&updatedPost, "id = ?", postId)
	if result.Error != nil {
		ResponseWithError(ctx, http.StatusNotFound, "No post with that title exists")
		return
	}
	postToUpdate := models.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Image:   payload.Image,
		UserID:  currentUser.ID,
	}

	pc.DB.Model(&updatedPost).Updates(postToUpdate)

	ResponseWithSuccess(ctx, http.StatusOK, updatedPost, "Post updated successfully")
}

func (pc *PostController) FindPostById(ctx *gin.Context) {
	postId := ctx.Param("postId")

	var post models.Post
	result := pc.DB.Preload("User").First(&post, "id = ?", postId)
	if result.Error != nil {
		ResponseWithError(ctx, http.StatusNotFound, "No post with that title exists")
		return
	}

	ResponseWithSuccess(ctx, http.StatusOK, post, "Post found successfully")
}

func (pc *PostController) FindPosts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []models.Post
	results := pc.DB.Limit(intLimit).Offset(offset).Preload("User").Find(&posts)
	if results.Error != nil {
		ResponseWithError(ctx, http.StatusBadGateway, results.Error.Error())
		return
	}

	ResponseWithSuccess(ctx, http.StatusOK, gin.H{
		"results": len(posts),
		"data":    posts,
	}, "Posts found successfully")
}

func (pc *PostController) DeletePost(ctx *gin.Context) {
	postId := ctx.Param("postId")

	result := pc.DB.Delete(&models.Post{}, "id = ?", postId)

	if result.Error != nil {
		ResponseWithError(ctx, http.StatusBadGateway, "No post with that title exists")
		return
	}

	ResponseWithSuccessWithoutData(ctx, http.StatusNoContent, "Post deleted successfully")
}
