package main

import (
	"net/http"

	engine "./engines"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Resp struct {
	Code    int64
	Message string
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check

			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})

	e.GET("/", func(c echo.Context) error {
		resp := &Resp{
			Code:    200,
			Message: "Transliterator API",
		}
		return c.JSON(http.StatusOK, resp)
	})
	e.GET("/transliterate", engine.MainHandler)
	e.POST("/upload", func(c echo.Context) error {
		resp := &Resp{
			Code:    200,
			Message: "Upload a file",
		}
		return c.JSON(http.StatusOK, resp)
	})

	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":3000"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
