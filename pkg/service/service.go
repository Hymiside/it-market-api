package service

import (
	"errors"
	"github.com/Hymiside/it-market-api/pkg/repository"
)

var (
	ErrTokenClaims = errors.New("token claims are not of type *tokenClaims")
	ErrParseJWT    = errors.New("error to parse jwt-token")
	ErrSignMethod  = errors.New("invalid signing method")
)

type authorization interface {
	ParseToken(token string) (string, error)
}

type Service struct {
	Auth authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: newAuthService(repo),
	}
}
