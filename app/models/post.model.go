package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title     string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	Image     string    `gorm:"not null" json:"image,omitempty"`
	UserID    uuid.UUID `json:"userId,omitempty"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty" `
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CreatePostRequest struct {
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image" binding:"required"`
}

type UpdatePost struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Image   string `json:"image,omitempty"`
}
