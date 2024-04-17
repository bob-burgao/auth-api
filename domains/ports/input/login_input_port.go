package domain_port_input

import (
	domains_models "auth/domains/models"
	"context"
)

type LoginInputPort interface {
	Login(ctx context.Context, login string, password string) (*domains_models.AuthResult, error)
}
