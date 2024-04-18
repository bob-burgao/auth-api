package controllers_v1

import (
	domain_config "auth/domains/config"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	logger zerolog.Logger
	envs   *domain_config.Environments
}

func NewUserController(envs *domain_config.Environments) *UserController {
	return &UserController{
		logger: log.With().Str("service", envs.MsName).Str("class", "UserController").Logger(),
		envs:   envs,
	}
}

func (c *UserController) Create(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}

func (c *UserController) Update(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}

func (c *UserController) Find(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}

func (c *UserController) Delete(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}
