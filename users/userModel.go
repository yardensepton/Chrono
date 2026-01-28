package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	LastName    string    `gorm:"not null" json:"lastName"`
	Email       string    `gorm:"uniqueIndex;not null" json:"email"`
	PhoneNumber string    `gorm:"not null" json:"phoneNumber"`
	Role        string    `gorm:"not null" json:"role"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func NewUser(req UserRequest) User {
	return User{
		ID:          uuid.New().String(),
		Name:        req.Name,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Role:        "user",
	}
}
