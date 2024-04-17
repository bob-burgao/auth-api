package repository_dynamo_mapper

import (
	repository_dynamo_entity "auth/adapters/repositories/dynamo/entities"
	domain_model "auth/domains/models"
)

func UserToDomainCustomer(user repository_dynamo_entity.User) *domain_model.Customer {
	return &domain_model.Customer{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Roles: user.Roles,
	}
}
