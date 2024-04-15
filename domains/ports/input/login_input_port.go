package domain_port_input

import domains_models "auth/domains/models"

type LoginInputPort interface {
	Login(login string, password string) (*domains_models.AuthResult, error)
}
