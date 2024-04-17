package domain_adapter

import (
	domain_model "auth/domains/models"
	domain_port_input "auth/domains/ports/input"
	domain_service "auth/domains/services"
	"context"
)

type AuthAdapter struct {
	authService domain_service.AuthService
}

func NewAuthAdapter(authService domain_service.AuthService) domain_port_input.LoginInputPort {
	return &AuthAdapter{
		authService: authService,
	}
}

type Item struct {
	Email    string
	Password string
}

func (a *AuthAdapter) Login(ctx context.Context, login string, password string) (*domain_model.AuthResult, error) {
	return a.authService.Login(ctx, login, password)
}
