package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jwhittle933/transliteratorAPI/controllers"
	mw "github.com/jwhittle933/transliteratorAPI/middleware"
	"github.com/jwhittle933/transliteratorAPI/types"

	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	/*
	 * sql.Open() doesn't directly open a conection
	 * and won't return error if the server isn't
	 * available or the conn data isn't correct.
	 *
	 * Thus, sql.Ping() is used to check for err
	 *
	 * MySQL driver imported in main
	 */
	db, _ := sql.Open("mysql", "root:[password]@tcp(127.0.0.1:3306)/transliterator")
	err := db.Ping()
	if err != nil {
		log.Panicln("Connected to DB.")
		log.Fatal(err)
	}
	defer db.Close()

	// MIDDLEWARE
	mw.MiddleWare(e)
	e.AutoTLSManager.Cache = autocert.DirCache("/cache/.cache")

	// ROUTES
	e.GET("/", baseRouteHandler)
	e.GET("/transliterate", controllers.TransliterateController)
	e.POST("/transliterate", controllers.TransliterateController)
	e.POST("/upload", controllers.UploadController)

	// !! USER ROUTES
	user := e.Group("/users")
	user.POST("/create", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		fmt.Println(email)
		fmt.Println(password)
		return c.JSON(http.StatusOK, "SIGNUP")
	})

	// !! SERVE STATIC FILES
	e.GET("/tmp/:user/:dir/:file", controllers.DownloadFile)

	// START
	e.Logger.Fatal(e.Start(":3000"))
}

// for dev purposes
func baseRouteHandler(c echo.Context) error {
	resp := &types.Resp{
		Code:    200,
		Message: "Transliterator API",
	}
	return c.JSON(http.StatusOK, resp)
}
