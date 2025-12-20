package users

type UserRequest struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Role        string `json:"role" binding:"required"`
}