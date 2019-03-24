package middlewarehelpers

import (
	"github.com/labstack/echo"
)

// Authenticate middleware.
func Authenticate(next echo.HandlerFunc, e echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the credentials from HTTP request header and perform a security
		// check

		// For invalid credentials
		// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

		// For valid credentials call next
		return next(c)
	}
}
