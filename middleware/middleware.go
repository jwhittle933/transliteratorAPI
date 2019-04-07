package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// MiddleWare func wrapper for setting all middleware
func MiddleWare(e *echo.Echo) {
	Logger(e)
	Authenticate(e)
	SetCors(e)
}

// SetCors for allowing Cross-Origin
func SetCors(e *echo.Echo) {
	e.Use(middleware.CORS())
}

// Authenticate middleware.
func Authenticate(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})
}

// Logger abstracts middleware logic. TODO
func Logger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}
