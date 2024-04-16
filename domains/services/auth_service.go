package domain_service

import (
	domain_model "auth/domains/models"
)

type AuthService struct {
	tokenService TokenService
}

func NewAuthService(tokenService TokenService) *AuthService {
	return &AuthService{
		tokenService: tokenService,
	}
}

func (a *AuthService) Login(login string, password string) (*domain_model.AuthResult, error) {

	customerData := domain_model.CustomerLogged{
		Id:    "068614e8-81fe-420a-a8a7-658ab0f5f706",
		Name:  "Albérico",
		Roles: []string{"commum", "teste"},
	}

	generatedToken, err := a.tokenService.GenerateToken(customerData)

	if err != nil {
		return nil, err
	}

	return generatedToken, nil
}
