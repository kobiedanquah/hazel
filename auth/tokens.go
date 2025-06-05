package auth

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/freekobie/hazel/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("the provided token is no valid")
)

type UserSession struct {
	User         models.User `json:"user"`
	RefreshToken string      `json:"refreshToken"`
	ExpiresAt    time.Time   `json:"expiresAt"`
}

func GenerateToken(userID uuid.UUID, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().UTC().UnixNano(),
		"exp": time.Now().Add(duration).UnixNano(),
		"sub": userID.String(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		slog.Error("failed to sign access token", "error", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return uuid.Nil, err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		slog.Error("failed to fetch 'sub' claim", "error", err.Error())
		return uuid.Nil, err
	}

	userID := uuid.MustParse(userIDString)

	return userID, nil
}
