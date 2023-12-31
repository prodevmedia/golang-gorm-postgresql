package fakers

import (
	"fmt"
	"os"

	"github.com/bxcodec/faker/v4"
	"github.com/supardi98/golang-gorm-postgres/app/models"
	"github.com/supardi98/golang-gorm-postgres/app/utils"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) models.User {
	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return models.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: hashedPassword,
		Role:     "user",
		Verified: true,
		Provider: "local",
	}
}
