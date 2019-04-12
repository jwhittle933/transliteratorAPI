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
	e, conn, err := start.Init()
	if err != nil {
		panic(err)
	}
	fmt.Println(conn)

	// MIDDLEWARE
	mw.MiddleWare(e)
	e.AutoTLSManager.Cache = autocert.DirCache("/cache/.cache")

	// ROUTES
	e.GET("/", baseRouteHandler)
	e.GET("/transliterate", controllers.TransliterateController)
	e.POST("/upload", controllers.UploadController)
	e.GET("/upload", uploadRouteHandler)

	// !! USER ROUTES
	user := e.Group("/users")
	user.GET("/signup", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "SIGNUP")
	})
	user.POST("/signup", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		fmt.Println(email)
		fmt.Println(password)
		return c.JSON(http.StatusOK, "SIGNUP")
	})

	// !! SERVE STATIC FILES
	e.GET("/tmp/:user/:dir/:file", func(c echo.Context) error {
		return c.File("/tmp/" + c.Param("file"))
	})

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

// for dev purposes
func uploadRouteHandler(c echo.Context) error {
	resp := &types.Resp{
		Code:    200,
		Message: "Upload a file",
	}
	return c.JSON(http.StatusOK, resp)
}
