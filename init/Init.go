package init

import "github.com/labstack/echo"

func Init() *echo.Echo {
	e := echo.New()
	return e
}
