package domain_service

import (
	domain_model "auth/domains/models"
	domain_port_output "auth/domains/ports/output"
	"context"
)

type AuthService struct {
	tokenService   TokenService
	userRepository domain_port_output.UserOutputPort
}

func NewAuthService(tokenService TokenService, userRepository domain_port_output.UserOutputPort) *AuthService {
	return &AuthService{
		tokenService:   tokenService,
		userRepository: userRepository,
	}
}

func (a *AuthService) Login(ctx context.Context, login string, password string) (*domain_model.AuthResult, error) {

	customerData, errFindUser := a.userRepository.FindUserByLoginAndPass(ctx, login, password)
	if errFindUser != nil {
		return nil, errFindUser
	}

	generatedToken, errGenerateToken := a.tokenService.GenerateToken(*customerData)
	if errGenerateToken != nil {
		return nil, errGenerateToken
	}

	return generatedToken, nil
}
