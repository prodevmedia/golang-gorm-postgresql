package seeders

import (
	"log"

	"github.com/prodevmedia/golang-gorm-postgres/app/models"
	"github.com/prodevmedia/golang-gorm-postgres/database/fakers"
	"gorm.io/gorm"
)

func PostSeeder(db *gorm.DB) {
	// Create
	// create array
	posts := []models.Post{}
	// use post faker
	for i := 0; i < 100; i++ {
		posts = append(posts, fakers.PostFaker(db))
	}

	// insert multiple record
	// mapping
	for _, post := range posts {
		db.Create(&post)
	}

	log.Println("? Seeding users completed")
}
