package repository_dynamo

import (
	repository_dynamo_entity "auth/adapters/repositories/dynamo/entities"
	repository_dynamo_mapper "auth/adapters/repositories/dynamo/mappers"
	domain_model "auth/domains/models"
	domain_port_output "auth/domains/ports/output"
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type UsersRepository struct {
	dataBase  *dynamodb.Client
	tableName string
}

func NewUsersRepository(dataBase *dynamodb.Client) domain_port_output.UserOutputPort {
	return &UsersRepository{
		dataBase:  dataBase,
		tableName: "users",
	}
}

func (u *UsersRepository) FindUserByLoginAndPass(ctx context.Context, login string, password string) (*domain_model.Customer, error) {

	attLogin, errLogin := attributevalue.Marshal(login)
	attPassword, errPass := attributevalue.Marshal(password)

	if errLogin != nil || errPass != nil {
		return nil, errors.New("error to convert input params")
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(u.tableName),
		Key: map[string]types.AttributeValue{
			"email": attLogin, "password": attPassword,
		},
	}
	result, err := u.dataBase.GetItem(ctx, input)

	if err != nil {
		return nil, errors.New("error to consult user by login and pass")
	}

	user := repository_dynamo_entity.User{}
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, errors.New("error to unmarshal result of consult by login and pass")
	}

	return repository_dynamo_mapper.UserToDomainCustomer(user), nil
}
