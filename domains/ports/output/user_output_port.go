package domain_port_output

import (
	domain_model "auth/domains/models"
	"context"
)

type UserOutputPort interface {
	FindUserByLoginAndPass(ctx context.Context, login string, password string) (*domain_model.Customer, error)
}
