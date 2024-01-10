package fakers

import (
	"github.com/bxcodec/faker/v4"
	"github.com/prodevmedia/golang-gorm-postgres/app/models"
	"gorm.io/gorm"
)

func PostFaker(db *gorm.DB) models.Post {
	// get user
	var user models.User
	// get random user
	db.First(&user, "role = ?", "user").Order("RANDOM()").Limit(1)

	return models.Post{
		Title:   faker.Name(),
		Content: faker.Paragraph(),
		Image:   "image.png",
		User:    user.ID,
	}
}
