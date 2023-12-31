package seeders

import (
	"log"

	"github.com/supardi98/golang-gorm-postgres/app/models"
	"github.com/supardi98/golang-gorm-postgres/database/fakers"
	"gorm.io/gorm"
)

func PostSeeder(db *gorm.DB) {
	// Create
	// create array
	posts := []models.Post{}
	// use user faker
	for i := 0; i < 100; i++ {
		posts = append(posts, fakers.PostFaker(db))
	}

	// insert multiple record
	// mapping
	for _, user := range posts {
		db.Create(&user)
	}

	log.Println("? Seeding users completed")
}
