package main

import (
	"encoding/json"
	"net/http"

	"database/sql"

	"./controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/thedevsaddam/govalidator"
)

// Resp struct for response schema.
type Resp struct {
	Code    int64
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"file:text": []string{"ext:txt, docx", "size:100000", "mime:txt, docx", "required"},
	}

	messages := govalidator.MapData{
		"file:text": []string{"ext:Only txt/docx allowed", "required:document is required"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	e := v.Validate()
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func main() {
	http.HandleFunc("/", handler)
	e := echo.New()
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

	e.GET("/", func(c echo.Context) error {
		resp := &Resp{
			Code:    200,
			Message: "Transliterator API",
		}
		return c.JSON(http.StatusOK, resp)
	}).Name = "home-route"

	e.GET("/transliterate", controllers.Transliterator).Name = "transliterate-query"
	e.GET("/upload", controllers.ProcessFile)
	e.POST("/upload", func(c echo.Context) error {
		resp := &Resp{
			Code:    200,
			Message: "Upload a file",
		}
		return c.JSON(http.StatusOK, resp)
	}).Name = "transliterate-upload"

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
