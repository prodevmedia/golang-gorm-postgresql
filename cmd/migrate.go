package cmd

import (
	"fmt"
	"log"

	"github.com/supardi98/golang-gorm-postgres/app/models"
	"github.com/supardi98/golang-gorm-postgres/config"
	"github.com/supardi98/golang-gorm-postgres/database"
)

func Migrate() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.ConnectDB(&config)

	database.DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("? Migration complete")
}
