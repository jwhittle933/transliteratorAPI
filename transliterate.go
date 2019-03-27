package main

import (
	"net/http"

	"database/sql"

	"./controllers"
	start "./init"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// https://github.com/golang/go/wiki/CodeReviewComments

// Resp struct for response schema.
type Resp struct {
	Code    int64
	Message string
}

func main() {
	e := start.Init()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})

	e.GET("/", baseRouteHandler)
	e.GET("/transliterate", controllers.Transliterator)
	e.POST("/upload", controllers.ProcessFile)
	e.GET("/upload", func(c echo.Context) error {
		resp := &Resp{
			Code:    200,
			Message: "Upload a file",
		}
		return c.JSON(http.StatusOK, resp)
	})

	e.GET("/users/:id", getUser)

	auth := e.Group("/auth")
	auth.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "Joe" && password == "password" {
			return true, nil
		}
		return false, nil
	}))

	e.Logger.Fatal(e.Start(":3000"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func initDb() (*sql.DB, error) {
	// https://github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql", "root:[password]@/transliterator")
	return db, err
}

func baseRouteHandler(c echo.Context) error {
	resp := &Resp{
		Code:    200,
		Message: "Transliterator API",
	}
	return c.JSON(http.StatusOK, resp)
}
