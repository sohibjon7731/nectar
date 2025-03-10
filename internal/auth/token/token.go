package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sohibjon7731/ecommerce_backend/config"
)

func GenerateTokens(userID uint) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Minute * 24).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %w", err)
	}

	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"sub": userID,
	// 	"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	// })

	// refreshTokenString, err := refreshToken.SignedString([]byte(config.JWTSecret))
	// if err != nil {
	// 	return "", "", fmt.Errorf("failed to create access token: %w", err)
	// }

	return accessTokenString, nil
}
