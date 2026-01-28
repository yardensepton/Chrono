package users

import (
	"time"
)

type RefreshToken struct {
	ID        string    `bson:"_id,omitempty"`
	UserID    string    `bson:"user_id"`
	TokenHash string    `bson:"token_hash" json:"token_hash" binding:"required"`
	ExpiresAt time.Time `bson:"expires_at"`
	CreatedAt time.Time `bson:"created_at"`
	UsedFlag  bool      `bson:"used_flag" json:"used_flag"`
}

func NewRefreshToken(userID string, jti string, tokenHash string, expiresAt time.Time) RefreshToken {
	return RefreshToken{
		ID:        jti,
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
		UsedFlag:  false,
	}
}
