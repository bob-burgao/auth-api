package routes

import (
	"auth/adapters/controllers"

	"github.com/labstack/echo/v4"
)

func SetUpHealthRoute(r *echo.Echo) {
	r.GET("/health", controllers.Health)
}
