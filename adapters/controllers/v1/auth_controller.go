package controllers_v1

import (
	domain_port_input "auth/domains/ports/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authAdapter domain_port_input.LoginInputPort
}

func NewLoginController(authAdapter domain_port_input.LoginInputPort) *AuthController {
	return &AuthController{
		authAdapter: authAdapter,
	}
}

func (c *AuthController) AuthWithLoginAndPass(ctx echo.Context) error {
	//body := ctx.Request().Body
	token, err := c.authAdapter.Login(ctx, "alberico", "123456")

	if err != nil {
		return ctx.String(http.StatusBadRequest, "Login error")
	}

	return ctx.JSON(http.StatusAccepted, token)
}
