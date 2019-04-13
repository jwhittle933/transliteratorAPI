package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jwhittle933/transliteratorAPI/controllers"
	start "github.com/jwhittle933/transliteratorAPI/init"
	mw "github.com/jwhittle933/transliteratorAPI/middleware"
	"github.com/jwhittle933/transliteratorAPI/types"

	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	app, err := start.Init()
	if err != nil {
		panic(err)
	}
	fmt.Println(app.DB)

	// MIDDLEWARE
	mw.MiddleWare(app.Echo)
	app.Echo.AutoTLSManager.Cache = autocert.DirCache("/cache/.cache")

	// ROUTES
	app.Echo.GET("/", baseRouteHandler)
	app.Echo.GET("/transliterate", controllers.TransliterateController)
	app.Echo.POST("/upload", controllers.UploadController)
	app.Echo.GET("/upload", uploadRouteHandler)

	// !! USER ROUTES
	user := app.Echo.Group("/users")
	user.POST("/create", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		fmt.Println(email)
		fmt.Println(password)
		return c.JSON(http.StatusOK, "SIGNUP")
	})

	// !! SERVE STATIC FILES
	app.Echo.GET("/tmp/:user/:dir/:file", controllers.DownloadFile)

	// START
	app.Echo.Logger.Fatal(app.Echo.Start(":3000"))
}

// for dev purposes
func baseRouteHandler(c echo.Context) error {
	resp := &types.Resp{
		Code:    200,
		Message: "Transliterator API",
	}
	return c.JSON(http.StatusOK, resp)
}

// for dev purposes
func uploadRouteHandler(c echo.Context) error {
	resp := &types.Resp{
		Code:    200,
		Message: "Upload a file",
	}
	return c.JSON(http.StatusOK, resp)
}
