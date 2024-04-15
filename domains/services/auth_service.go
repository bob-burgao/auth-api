package domain_service

import (
	domain_model "auth/domains/models"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(login string, password string) (*domain_model.AuthResult, error) {
	panic("unimplemented")
}
