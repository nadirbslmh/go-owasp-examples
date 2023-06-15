package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// CORS configuration
	config := middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}

	e := echo.New()

	e.Use(middleware.Logger())

	// use CORS for all routes
	e.Use(middleware.CORSWithConfig(config))

	e.GET("/resources", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is resources")
	})

	e.POST("/create", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "create sample")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
