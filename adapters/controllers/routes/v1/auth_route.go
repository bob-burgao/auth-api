package routes_v1

import (
	controllers_v1 "auth/adapters/controllers/v1"

	"github.com/labstack/echo/v4"
)

func SetUpAuthRoute(r *echo.Echo) {
	basePath := "/api/v1/auth"

	authController := controllers_v1.NewLoginController()

	r.POST(basePath, authController.AuthWithLoginAndPass)
}
