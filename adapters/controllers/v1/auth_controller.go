package controllers_v1

import (
	adapter_util "auth/adapters/utils"
	domain_config "auth/domains/config"
	domain_port_input "auth/domains/ports/input"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AuthController struct {
	authAdapter domain_port_input.LoginInputPort
	logger      zerolog.Logger
	envs        *domain_config.Environments
}

func NewLoginController(authAdapter domain_port_input.LoginInputPort, envs *domain_config.Environments) *AuthController {

	return &AuthController{
		authAdapter: authAdapter,
		logger:      log.With().Str("service", envs.MsName).Str("class", "AuthController").Logger(),
		envs:        envs,
	}
}

func (c *AuthController) AuthWithLoginAndPass(ctx echo.Context) error {
	requiredParams := []string{"login", "password"}
	jsonBody, errGetParams := adapter_util.GetControllerBodyData(ctx.Request(), requiredParams)
	if errGetParams != nil {
		return ctx.String(http.StatusBadRequest, errGetParams.Error())
	}

	body := *jsonBody
	login := body[requiredParams[0]].(string)
	pass := body[requiredParams[1]].(string)

	c.logger.Info().Msg(fmt.Sprintf("init login: %s", login))

	token, err := c.authAdapter.Login(ctx.Request().Context(), login, pass)

	if err != nil {
		return ctx.String(http.StatusUnauthorized, "Login error")
	}

	c.logger.Info().Msg(fmt.Sprintf("success login: %s", login))

	return ctx.JSON(http.StatusAccepted, token)
}

func (c *AuthController) RecoverPass(ctx echo.Context) error {
	// TODO - Não implementado
	return ctx.String(http.StatusAccepted, "feito")
}

func (c *AuthController) ChangePass(ctx echo.Context) error {
	// TODO - Não implementado
	return ctx.String(http.StatusAccepted, "feito")
}
