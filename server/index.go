package server

import (
	"github.com/labstack/echo"
)

func setupMiddleware() *echo.Echo {
	e := echo.New()

	// set all middlewares
	middleware.setLogMiddleware(e)

	return e
}
