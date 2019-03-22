package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Text struct {
	Body string `json:"body" form:"body" query:"body"`
	Lang string `json:"lang" form:"lang" query:"lang"`
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
		return c.JSON(http.StatusOK, `{"code": 200, "message": "transliterator API" }`)
	})

	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	})
	e.GET("/users/:id", getUser)
	e.GET("/hebrew", hebrew)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func hebrew(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func NewText(c echo.Context) (err error) {
	t := new(Text)
	if err = c.Bind(t); err != nil {
		return
	}
	return c.JSON(http.StatusOK, t)
}
