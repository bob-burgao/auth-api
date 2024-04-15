package domain_adapter

import (
	domain_model "auth/domains/models"
	domain_port_input "auth/domains/ports/input"
	domain_service "auth/domains/services"
)

type AuthAdapter struct {
	authService domain_service.AuthService
}

func NewAuthAdapter(authService domain_service.AuthService) domain_port_input.LoginInputPort {
	return &AuthAdapter{
		authService: authService,
	}
}

func (a *AuthAdapter) Login(login string, password string) (*domain_model.AuthResult, error) {
	return a.authService.Login(login, password)
}
