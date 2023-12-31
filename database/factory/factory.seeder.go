package factory

import (
	"github.com/supardi98/golang-gorm-postgres/app/models"
	"gorm.io/gorm"
)

func UserFactory(db *gorm.DB) models.User {
	return models.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: "asdas",
		Role:     "user",
		Verified: true,
		Provider: "local",
	}
}
