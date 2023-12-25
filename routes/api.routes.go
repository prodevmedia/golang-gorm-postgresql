package routes

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supardi98/golang-gorm-postgres/config"
	"github.com/supardi98/golang-gorm-postgres/database"
	"gorm.io/gorm"
)

func InitRoutes() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.ConnectDB(&config)

	dbConnection := database.DB

	NewRoute(dbConnection, config)
}

func NewRoute(dbConnection *gorm.DB, config config.Config) {
	var (
		AuthRouteController = NewAuthRouteController(database.DB)
		UserRouteController = NewRouteUserController(database.DB)
		PostRouteController = NewRoutePostController(database.DB)
	)

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}