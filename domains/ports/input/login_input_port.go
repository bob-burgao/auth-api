package domain_port_input

import (
	domains_models "auth/domains/models"

	"github.com/labstack/echo/v4"
)

type LoginInputPort interface {
	Login(ctx echo.Context, login string, password string) (*domains_models.AuthResult, error)
}
