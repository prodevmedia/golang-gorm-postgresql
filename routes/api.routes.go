package routes

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/prodevmedia/golang-gorm-postgres/config"
	"github.com/prodevmedia/golang-gorm-postgres/database"
	"gorm.io/gorm"
)

func ApiRoute(router *gin.RouterGroup, dbConnection *gorm.DB, config config.Config) {
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// Public routes
	router.Static("/public", "./public")

	WSRoute(router, database.DB)

	AuthRoute(router, database.DB)
	ProfileRoute(router, database.DB)
	UserRoute(router, database.DB)
	PostRoute(router, database.DB)
}

func Init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// OAUTH PROVIDERS
	goth.UseProviders(
		google.New(config.OAuthGoogleClientId, config.OAuthGoogleSecret, config.ServerURL+"/auth/provider/google/callback", "email", "profile"),
		github.New(config.OAuthGithubClientId, config.OAuthGithubSecret, config.ServerURL+"/auth/provider/github/callback", "user", "gist", "user:email"),
	)

	database.ConnectDB(&config)

	// server := gin.Default()
	server := gin.New()
	server.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:" + config.ServerPort, config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	// router := server.Group("/api")
	router := server.Group("/")
	ApiRoute(router, database.DB, config)

	log.Fatal(server.Run(":" + config.ServerPort))
}
