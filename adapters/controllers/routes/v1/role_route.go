package routes_v1

import (
	controllers_v1 "auth/adapters/controllers/v1"
	domain_config "auth/domains/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func SetRoleUserRoute(r *echo.Echo, db *dynamodb.Client) {
	envs := domain_config.LoadEnvVars()
	path := envs.BasePathV1 + "/role"

	roleController := controllers_v1.NewRoleController(envs)

	r.GET(path, roleController.Find)
	r.POST(path, roleController.Create)
	r.PUT(path, roleController.Update)
	r.DELETE(path, roleController.Delete)
}
