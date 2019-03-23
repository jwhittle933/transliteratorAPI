package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Logger() {
	middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})
}

func Authenticate(next echo.HandlerFunc, e echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the credentials from HTTP request header and perform a security
		// check

		// For invalid credentials
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

		// For valid credentials call next
		// return next(c)
	}
}