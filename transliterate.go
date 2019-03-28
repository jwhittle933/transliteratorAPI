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

var middlewareConfig = middleware.LoggerConfig{
	Format: "method=${method}, uri=${uri}, status=${status}\n",
}

func main() {
	e := start.Init()
	e.Use(middleware.LoggerWithConfig(middlewareConfig))
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
	e.POST("/upload", controllers.Uploader)
	http.HandleFunc("upload", controllers.UploadHandler)
	e.GET("/upload", uploadRouteHandler)
	e.Static("/tmp", "tmp")

	auth := e.Group("/auth")
	auth.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "Joe" && password == "password" {
			return true, nil
		}
		return false, nil
	}))

	e.Logger.Fatal(e.Start(":3000"))
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDb() (*sql.DB, error) {
	// https://github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql", "root:[password]@/transliterator")
	return db, err
}

func baseRouteHandler(c echo.Context) error {
	resp := &controllers.Resp{
		Code:    200,
		Message: "Transliterator API",
	}
	return c.JSON(http.StatusOK, resp)
}

func uploadRouteHandler(c echo.Context) error {
	resp := &controllers.Resp{
		Code:    200,
		Message: "Upload a file",
	}
	return c.JSON(http.StatusOK, resp)
}
