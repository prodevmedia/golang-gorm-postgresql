package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title     string    `gorm:"uniqueIndex;not null"`
	Content   string    `gorm:"not null"`
	Image     string    `gorm:"not null"`
	UserID    uuid.UUID `gorm:"not null;type:uuid"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePostRequest struct {
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image" binding:"required"`
	User    string `json:"user,omitempty"`
}

type UpdatePost struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Image   string `json:"image,omitempty"`
	User    string `json:"user,omitempty"`
}
