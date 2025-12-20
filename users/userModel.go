package users

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Role        string `json:"role"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func NewUser(req UserRequest) User {
	return User{
		ID:          uuid.New().String(),
		Name:        req.Name,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Role:        "user",
		CreatedAt:   time.Now().Format(time.RFC3339),
	}
}
