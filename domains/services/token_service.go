package domain_service

import (
	domain_model "auth/domains/models"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (a *TokenService) GenerateToken(customer domain_model.CustomerLogged) (*domain_model.AuthResult, error) {

	now := time.Now()
	timeToExpire := now.Add(1 * time.Hour * 24)
	expiresAt := jwt.NewNumericDate(timeToExpire)
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		&jwt.RegisteredClaims{
			Audience:  customer.Roles,
			ExpiresAt: expiresAt,
			ID:        customer.Id,
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "sso.example.com",
			NotBefore: jwt.NewNumericDate(now),
			Subject:   customer.Name,
		},
	)

	tokenSigned, err := token.SignedString(key)

	if err != nil {
		return nil, err
	}

	return &domain_model.AuthResult{
		Token:      tokenSigned,
		ExpireTime: timeToExpire.Second(),
	}, nil
}
