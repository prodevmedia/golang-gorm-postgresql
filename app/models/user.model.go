package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name               string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email              string    `gorm:"uniqueIndex;not null" json:"email,omitempty"`
	Password           string    `gorm:"not null" json:"-"`
	Role               string    `gorm:"type:varchar(255);not null" json:"role,omitempty"`
	Provider           string    `gorm:"not null" json:"provider,omitempty"`
	Avatar             string    `json:"avatar,omitempty"`
	VerificationCode   string    `json:"-"`
	PasswordResetToken string    `json:"-"`
	PasswordResetAt    time.Time `json:"passwordResetAt,omitempty"`
	Verified           bool      `gorm:"not null" json:"verified,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	Posts              []Post    `gorm:"foreignKey:UserID" json:"posts,omitempty"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// ? ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// ? UpdatePasswordInput struct
type UpdatePasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type UpdateProfileInput struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
