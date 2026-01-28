package services

import (
	"my-go-server/model/users"
	"my-go-server/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) InsertUser(userReq users.UserRequest) (users.User, error) {
	err := userReq.ValidateUserRequest()
	if err != nil {
		return users.User{}, err
	}
	userModel, err := users.NewUser(userReq)
	if err != nil {
		return users.User{}, err
	}
	return s.repo.Insert(userModel)
}

func (s *UserService) GetUserByID(id string) (users.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(user users.User) (users.User, error) {
	return s.repo.Update(user)
}
func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

// // AuthenticateUser authenticates a user with email and password
// func (s *UserService) AuthenticateUser(email, password string) (users.User, string, string, error) {
// 	user, err := s.GetUserByEmail(email)
// 	if err != nil {
// 		return users.User{}, "", "", errors.New("invalid credentials")
// 	}

// 	// Check password
// 	if !user.CheckPassword(password) {
// 		return users.User{}, "", "", errors.New("invalid credentials")
// 	}

// 	// Generate JWT tokens
// 	accessToken, err := utils.GenerateAccessToken(user)
// 	if err != nil {
// 		return users.User{}, "", "", err
// 	}
// 	refreshToken, err := utils.GenerateRefreshToken(user)
// 	if err != nil {
// 		return users.User{}, "", "", err
// 	}

// 	// Store refresh token in the database
// 	utils.HashToken(refreshToken)
// 	refreshTokenModel := users.NewRefreshToken(user.ID, refreshToken, time.Now().Add(24*time.Hour))
// 	_, err = s.refreshTokenRepo.Insert(refreshTokenModel)
// 	if err != nil {
// 		return users.User{}, "", "", err
// 	}

// 	return user, accessToken, refreshToken, nil
// }

// GetUserByEmail finds a user by email address
func (s *UserService) GetUserByEmail(email string) (users.User, error) {
	return s.repo.GetByEmail(email)
}
