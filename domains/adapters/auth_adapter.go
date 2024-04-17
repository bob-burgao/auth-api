package domain_adapter

import (
	domain_model "auth/domains/models"
	domain_port_input "auth/domains/ports/input"
	domain_service "auth/domains/services"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

type AuthAdapter struct {
	authService domain_service.AuthService
	dataBase    *dynamodb.Client
}

func NewAuthAdapter(dataBase *dynamodb.Client, authService domain_service.AuthService) domain_port_input.LoginInputPort {
	return &AuthAdapter{
		authService: authService,
		dataBase:    dataBase,
	}
}

type Item struct {
	Email    string
	Password string
}

func (a *AuthAdapter) Login(ctx echo.Context, login string, password string) (*domain_model.AuthResult, error) {

	tableName := "users"

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := a.dataBase.Scan(ctx.Request().Context(), input)

	if err != nil {
		fmt.Println("error get data dynamo")
	} else {
		fmt.Println("funcinou get data dynamo", result)
	}

	for _, i := range result.Items {
		item := Item{}
		err = attributevalue.UnmarshalMap(i, &item)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}

		fmt.Println("Email: ", item.Email)
		fmt.Println("Pass:", item.Password)
	}

	return a.authService.Login(login, password)
}
