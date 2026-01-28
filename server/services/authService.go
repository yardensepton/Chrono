package services

import (
	"errors"
	"log"
	"my-go-server/model/users"
	"my-go-server/repositories"
	"my-go-server/utils"
	"os"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo repositories.UserRepository
	refreshTokenRepo repositories.RefreshTokenRepository
}

func NewAuthService(userRepo repositories.UserRepository, refreshTokenRepo repositories.RefreshTokenRepository) *AuthService {
	return &AuthService{userRepo: userRepo, refreshTokenRepo: refreshTokenRepo}
}

func getAccessTokenExpireMinutes() time.Duration {
	minutesStr := os.Getenv("ACCESS_TOKEN_EXPIRE_MINUTES")
	if minutesStr == "" {
		return 60 // default 60 minutes
	}
	minutes, err := time.ParseDuration(minutesStr + "m")
	if err != nil {
		return 60
	}
	return minutes
}

func getRefreshTokenExpireMinutes() time.Duration {
	minutesStr := os.Getenv("REFRESH_TOKEN_EXPIRE_MINUTES")
	if minutesStr == "" {
		return 2880 // default 2 days
	}
	minutes, err := time.ParseDuration(minutesStr + "m")
	if err != nil {
		return 2880
	}
	return minutes
}

func(s *AuthService) GennerateTokens(user users.User) (string, string, error) {
	accessTokenExpireMinutes := getAccessTokenExpireMinutes()
	refreshTokenExpireMinutes := getRefreshTokenExpireMinutes()
	
	// Generate JWT tokens
	accessTokenJti := uuid.New().String()
	accessToken, err := utils.GenerateAccessToken(user, accessTokenExpireMinutes, accessTokenJti)
	if err != nil {		
		return  "", "", errors.New("access token generation failed")
	}
	refreshTokenJti := uuid.New().String()
	refreshToken, err := utils.GenerateRefreshToken(user, refreshTokenExpireMinutes, refreshTokenJti)
	if err != nil {
		return "", "", errors.New("refresh token generation failed")
	}
	expired := time.Now().Add(refreshTokenExpireMinutes * time.Minute)
	log.Println(expired)
	hashedRefreshToken := utils.HashToken(refreshToken)
	refreshTokenModel := users.NewRefreshToken(user.ID, refreshTokenJti, hashedRefreshToken, expired)
		_, err = s.refreshTokenRepo.Insert(refreshTokenModel)
	if err != nil {
		return "", "", errors.New("failed to store refresh token")
	}
	return accessToken, refreshToken, nil
}

// AuthenticateUser authenticates a user with email and password
func (s *AuthService) AuthenticateUser(email, password string) (users.User, string, string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return users.User{}, "", "", errors.New("invalid credentials")
	}

	// Check password
	if !user.CheckPassword(password) {
		return users.User{}, "", "", errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := s.GennerateTokens(user)
	if err != nil {
		return users.User{}, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

// RefreshTokens refreshes access and refresh tokens
func (s *AuthService) RefreshTokens(oldRefreshToken string) (string, string, error) {
	hashedToken := utils.HashToken(oldRefreshToken)
	storedToken, err := s.refreshTokenRepo.GetByToken(hashedToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	if storedToken.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("refresh token expired")
	}

	if storedToken.UsedFlag {
		return "", "", errors.New("refresh token already used")
	}

	user, err := s.userRepo.GetByID(storedToken.UserID)
	if err != nil {
		return "", "", errors.New("user not found")
	}
	newAccessToken, newRefreshToken, err := s.GennerateTokens(user)
	if err != nil {
		return "", "", err
	}

	// Update stored refresh token
	storedToken.UsedFlag = true

	storedToken, err = s.refreshTokenRepo.Update(storedToken)
	if err != nil {
		return "", "", errors.New("failed to update old refresh token")
	}

	return newAccessToken, newRefreshToken, nil
}