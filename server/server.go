package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MainGetTest(ctx echo.Context) error {
	print(ctx)

	return ctx.String(http.StatusOK, "Success")
}
