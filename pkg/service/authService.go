package service

import (
	"time"

	"github.com/Hymiside/it-market-api/pkg/repository"
	"github.com/golang-jwt/jwt"
)

var (
	signingKey = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")
	tokenTTL   = 1460 * time.Hour
)

type Claims struct {
	jwt.StandardClaims
	UserId string
}

type authService struct {
	repo *repository.Repository
}

func newAuthService(repo *repository.Repository) *authService {
	return &authService{repo: repo}
}

func (a *authService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSignMethod
		}
		return signingKey, nil
	})
	if err != nil {
		return "", ErrParseJWT
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", ErrTokenClaims
	}
	return claims.UserId, nil
}
