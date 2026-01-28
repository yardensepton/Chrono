package users

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          string              `json:"id" bson:"_id,omitempty"`
	Name        string              `json:"name" bson:"name"`
	LastName    string              `json:"lastName" bson:"lastName"`
	Email       string              `json:"email" bson:"email"`
	Password    string              `json:"-" bson:"password"` // Don't serialize password in JSON
	PhoneNumber string              `json:"phoneNumber" bson:"phoneNumber"`
	Role        Role                `json:"role" bson:"role"`
	TherapistID *primitive.ObjectID `bson:"therapist_id,omitempty"` //null for therapists
	Permissions []string            `bson:"permissions,omitempty"`
	CreatedAt   time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt" bson:"updatedAt"`
}

func NewUser(req UserRequest) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:          uuid.New().String(),
		Name:        req.Name,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    string(hashedPassword),
		PhoneNumber: req.PhoneNumber,
		Role:        Role(req.Role),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// CheckPassword verifies if the provided password matches the hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
