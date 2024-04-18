package controllers_v1

import (
	domain_config "auth/domains/config"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RoleController struct {
	logger zerolog.Logger
	envs   *domain_config.Environments
}

func NewRoleController(envs *domain_config.Environments) *RoleController {
	return &RoleController{
		logger: log.With().Str("service", envs.MsName).Str("class", "RoleController").Logger(),
		envs:   envs,
	}
}

func (c *RoleController) Create(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}

func (c *RoleController) Update(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}

func (c *RoleController) Find(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}

func (c *RoleController) Delete(ctx echo.Context) error {
	// TODO - Não implementado
	panic("not implemented")
}
