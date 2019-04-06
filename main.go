package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := Init()

	// MIDDLEWARE
	e.Use(middleware.LoggerWithConfig(middlewareConfig))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})

	// ROUTES
	e.GET("/", baseRouteHandler)
	e.GET("/transliterate", TransliterateController)
	e.POST("/upload", UploadController)
	e.GET("/upload", uploadRouteHandler)
	e.Static("/tmp", "tmp")

	auth := e.Group("/auth")
	auth.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "Joe" && password == "password" {
			return true, nil
		}
		return false, nil
	}))

	// START
	e.Logger.Fatal(e.Start(":3000"))
}

func baseRouteHandler(c echo.Context) error {
	resp := &Resp{
		Code:    200,
		Message: "Transliterator API",
	}
	return c.JSON(http.StatusOK, resp)
}

func uploadRouteHandler(c echo.Context) error {
	resp := &Resp{
		Code:    200,
		Message: "Upload a file",
	}
	return c.JSON(http.StatusOK, resp)
}
