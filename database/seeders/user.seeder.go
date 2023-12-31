package seeders

import (
	"log"

	"github.com/supardi98/golang-gorm-postgres/app/models"
	"github.com/supardi98/golang-gorm-postgres/app/utils"
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

	// insert multiple record
	// mapping
	for _, user := range users {
		db.Create(&user)
	}

	log.Println("? Seeding users completed")
}
