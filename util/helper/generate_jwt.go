package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenClaims
type TokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken help generate a token with secretKey
func GenerateToken(secretKey string, clams TokenClaims) (string, error) {
	clams.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 4230).Unix(),
	}
	bytes := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)
	return token.SignedString(bytes)
}
