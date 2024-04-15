package controllers_v1

import (
	domain_adapter "auth/domains/adapters"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authAdapter domain_adapter.AuthAdapter
}

func NewLoginController(authAdapter domain_adapter.AuthAdapter) *AuthController {
	return &AuthController{
		authAdapter: authAdapter,
	}
}

func (c *AuthController) AuthWithLoginAndPass(ctx echo.Context) error {
	//body := ctx.Request().Body
	c.authAdapter.Login("alberico", "123456")
	return ctx.String(http.StatusProcessing, "Recived the request with success, will be working to make the best report for you!!!")
}
