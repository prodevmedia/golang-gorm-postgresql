package cmd

import (
	"log"

	"github.com/prodevmedia/golang-gorm-postgres/config"
	"github.com/prodevmedia/golang-gorm-postgres/database"
	"github.com/prodevmedia/golang-gorm-postgres/database/seeders"
)

func Seed() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.ConnectDB(&config)

	seeders.UserSeeder(database.DB)
	seeders.PostSeeder(database.DB)
	log.Println("? Seeding completed")
}
