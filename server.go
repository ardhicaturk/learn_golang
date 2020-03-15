package main

import (
	"github.com/ardhicaturk/learn_golang/routes"

	"github.com/ardhicaturk/learn_golang/webserver/setupmiddleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define custom middleware
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

	// set group API by version
	apiV1 := e.Group("/api/v1")
	routes.V1(apiV1)

	// Server init on port 3233
	e.Logger.Fatal(e.Start(":3233"))
}
