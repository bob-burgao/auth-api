package domain_service

import (
	domain_config "auth/domains/config"
	domain_model "auth/domains/models"
	domain_port_output "auth/domains/ports/output"
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AuthService struct {
	tokenService   TokenService
	userRepository domain_port_output.UserOutputPort
	logger         zerolog.Logger
	envs           *domain_config.Environments
}

func NewAuthService(tokenService TokenService, userRepository domain_port_output.UserOutputPort, envs *domain_config.Environments) *AuthService {
	return &AuthService{
		tokenService:   tokenService,
		userRepository: userRepository,
		logger:         log.With().Str("service", envs.MsName).Str("class", "AuthService").Logger(),
		envs:           envs,
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
