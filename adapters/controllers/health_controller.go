package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "App is healthy")
}
