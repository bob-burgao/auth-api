package controllers_v1

import (
	domain_port_input "auth/domains/ports/input"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

	login := jsonBody["login"].(string)
	pass := jsonBody["password"].(string)

	token, err := c.authAdapter.Login(ctx.Request().Context(), login, pass)

	if err != nil {
		return ctx.String(http.StatusBadRequest, "Login error")
	}

	return ctx.JSON(http.StatusAccepted, token)
}
