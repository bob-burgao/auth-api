package routes_v1

import (
	controllers_v1 "auth/adapters/controllers/v1"
	domain_config "auth/domains/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func SetUpUserRoute(r *echo.Echo, db *dynamodb.Client) {
	envs := domain_config.LoadEnvVars()
	path := envs.BasePathV1 + "/user"

	userController := controllers_v1.NewUserController(envs)

	r.GET(path, userController.Find)
	r.POST(path, userController.Create)
	r.PUT(path, userController.Update)
	r.DELETE(path, userController.Delete)
}
