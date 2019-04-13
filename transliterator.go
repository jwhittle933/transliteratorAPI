package main

import (
	"fmt"
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
