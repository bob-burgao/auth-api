package repository_dynamo

import (
	repository_dynamo_entity "auth/adapters/repositories/dynamo/entities"
	repository_dynamo_mapper "auth/adapters/repositories/dynamo/mappers"
	domain_config "auth/domains/config"
	domain_model "auth/domains/models"
	domain_port_output "auth/domains/ports/output"
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type UsersRepository struct {
	dataBase  *dynamodb.Client
	tableName string
	logger    zerolog.Logger
	envs      *domain_config.Environments
}

func NewUsersRepository(dataBase *dynamodb.Client, envs *domain_config.Environments) domain_port_output.UserOutputPort {
	return &UsersRepository{
		dataBase:  dataBase,
		tableName: "users",
		logger:    log.With().Str("service", envs.MsName).Str("class", "UsersRepository").Logger(),
		envs:      envs,
	}
}

func (u *UsersRepository) FindUserByLoginAndPass(ctx context.Context, login string, password string) (*domain_model.Customer, error) {

	attLogin, errLogin := attributevalue.Marshal(login)
	attPassword, errPass := attributevalue.Marshal(password)

	if errLogin != nil || errPass != nil {
		message := fmt.Sprintf("error to convert input params: %s", login)
		u.logger.Error().Msg(message)
		return nil, errors.New(message)
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(u.tableName),
		Key: map[string]types.AttributeValue{
			"email": attLogin, "password": attPassword,
		},
	}
	result, err := u.dataBase.GetItem(ctx, input)

	if err != nil {
		message := fmt.Sprintf("error to consult user by login and pass: %s", login)
		u.logger.Error().Msg(message)
		return nil, errors.New(message)
	}

	user := repository_dynamo_entity.User{}
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		message := fmt.Sprintf("error to unmarshal result of consult by login and pass: %s", login)
		u.logger.Error().Msg(message)
		return nil, errors.New(message)
	}

	return repository_dynamo_mapper.UserToDomainCustomer(user), nil
}
