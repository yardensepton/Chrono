package utils

import (
	"log"
	"os"
	"time"
	"my-go-server/model/users"

	"github.com/golang-jwt/jwt/v5"
)

// var jwtSecret = getJWTSecret()
// var accessTokenExpireMinutes = getAccessTokenExpireMinutes()
// var refreshTokenExpireMinutes = getRefreshTokenExpireMinutes()

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not set")
	}
	return []byte(secret)

}


type Claims struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// generates a JWT access token for a user
func GenerateAccessToken(user users.User, ttl time.Duration, jti string) (string, error) {
	expirationTime := time.Now().Add(ttl) // Token expires in 24 hours

	claims := &Claims{
		ID : jti,
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-go-server",
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
// generates a JWT refresh token for a user
func GenerateRefreshToken(user users.User, ttl time.Duration, jti string) (string, error) {
	expirationTime := time.Now().Add(ttl) // Token expires in 24 hours

	claims := &Claims{
		ID:     jti,
		UserID: user.ID,
		Email:  user.Email,
		Role:   string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-go-server",
			Subject:   user.ID,
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := refreshToken.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
