package routes_v1

import (
	controllers_v1 "auth/adapters/controllers/v1"
	repository_dynamo "auth/adapters/repositories/dynamo"
	domain_adapter "auth/domains/adapters"
	domain_service "auth/domains/services"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func SetUpAuthRoute(r *echo.Echo, db *dynamodb.Client) {
	basePath := "/api/v1/auth"

	userRepository := repository_dynamo.NewUsersRepository(db)
	tokenService := domain_service.NewTokenService()
	authService := domain_service.NewAuthService(*tokenService, userRepository)
	authAdapter := domain_adapter.NewAuthAdapter(*authService)

	authController := controllers_v1.NewLoginController(authAdapter)

	r.POST(basePath, authController.AuthWithLoginAndPass)
}
