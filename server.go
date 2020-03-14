package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ardhicaturk/learn_golang/webserver/setupmiddleware"
)

func main() {
	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	setupmiddleware.SetLogMiddleware(e)

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
		},
	}))

	// router & handler for "/"
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok\n")
	})

	// Server
	e.Logger.Fatal(e.Start(":3233"))
}
