package routes_v1

import (
	controllers_v1 "auth/adapters/controllers/v1"
	repository_dynamo "auth/adapters/repositories/dynamo"
	domain_adapter "auth/domains/adapters"
	domain_config "auth/domains/config"
	domain_service "auth/domains/services"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func SetUpAuthRoute(r *echo.Echo, db *dynamodb.Client) {
	envs := domain_config.LoadEnvVars()

	//Init Data Repository
	userRepository := repository_dynamo.NewUsersRepository(db, envs)

	//Init Domain Services
	tokenService := domain_service.NewTokenService()
	authService := domain_service.NewAuthService(*tokenService, userRepository, envs)

	//Init Domain Input
	authAdapter := domain_adapter.NewAuthAdapter(*authService)

	//Init controller
	authController := controllers_v1.NewLoginController(authAdapter, envs)

	r.POST(envs.BasePathV1+"/auth", authController.AuthWithLoginAndPass)
	r.POST(envs.BasePathV1+"/recover-pass", authController.RecoverPass)
	r.POST(envs.BasePathV1+"/change-pass", authController.ChangePass)
}
