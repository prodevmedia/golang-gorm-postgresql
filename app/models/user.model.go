package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name               string    `gorm:"type:varchar(255);not null"`
	Email              string    `gorm:"uniqueIndex;not null"`
	Password           string    `gorm:"not null"`
	Role               string    `gorm:"type:varchar(255);not null"`
	Provider           string    `gorm:"not null"`
	Avatar             string
	VerificationCode   string
	PasswordResetToken string
	PasswordResetAt    time.Time
	Verified           bool `gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Posts              []Post `gorm:"foreignKey:UserID"`
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

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ? ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// ? UpdatePasswordInput struct
type UpdatePasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type UpdateProfileInput struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// User to UserResponse
func (u *User) Response() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Avatar:    u.Avatar,
		Provider:  u.Provider,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
