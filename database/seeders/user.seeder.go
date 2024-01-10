package seeders

import (
	"log"

	"github.com/prodevmedia/golang-gorm-postgres/app/models"
	"github.com/prodevmedia/golang-gorm-postgres/app/utils"
	"github.com/prodevmedia/golang-gorm-postgres/database/fakers"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	// Create
	// create array
	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		return
	}

	var users = []models.User{
		{
			Name:     "Admin",
			Email:    "admin@gmail.com",
			Password: hashedPassword,
			Role:     "user",
			Verified: true,
			Provider: "local",
		},
	}

	// use user faker
	for i := 0; i < 1000; i++ {
		users = append(users, fakers.UserFaker(db))
	}

	// insert multiple record
	// mapping
	for _, user := range users {
		db.Create(&user)
	}

	log.Println("? Seeding users completed")
}
