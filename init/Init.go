package start

import "github.com/labstack/echo"

// Init func for instantiating
func Init() *echo.Echo {
	e := echo.New()
	return e
}
