package users

import "fmt"

type UserRequest struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6" example:"secret123"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Role        Role   `json:"role" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (req UserRequest) ValidateUserRequest() error {
	if req.Role != RoleTherapist && req.Role != RoleAssistant && req.Role != RolePatient {
		return fmt.Errorf("invalid role: %s", req.Role)
	}
	return nil
}