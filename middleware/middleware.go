package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// MiddleWare func wrapper for setting all middleware
func MiddleWare(e *echo.Echo) {
	logger(e)
	recover(e)
	authenticate(e)
	sendJWT(e)
}

func logger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}

func setCors(e *echo.Echo) {
	e.Use(middleware.CORS())
}

func recover(e *echo.Echo) {
	e.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
}

func authenticate(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})
}

func sendJWT(e *echo.Echo) {
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("transliteratedToken"),
		TokenLookup: "query:token",
	}))
}
