package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supardi98/golang-gorm-postgres/app/controllers"
	"github.com/supardi98/golang-gorm-postgres/app/middleware"
	"gorm.io/gorm"
)

func PostRoute(rg *gin.RouterGroup, dbConnection *gorm.DB) {
	postController := controllers.NewPostController(dbConnection)

	router := rg.Group("posts")
	router.Use(middleware.DeserializeUser())
	router.POST("/", postController.CreatePost)
	router.GET("/", postController.FindPosts)
	router.PUT("/:postId", postController.UpdatePost)
	router.GET("/:postId", postController.FindPostById)
	router.DELETE("/:postId", postController.DeletePost)

}
