package main

import (
	"net/http"

	"github.com/jwhittle933/transliteratorAPI/controllers"
	start "github.com/jwhittle933/transliteratorAPI/init"
	mw "github.com/jwhittle933/transliteratorAPI/middleware"
	"github.com/jwhittle933/transliteratorAPI/types"

	"github.com/labstack/echo"
)

func main() {
	e := start.Init()

	// MIDDLEWARE
	mw.MiddleWare(e)

	// ROUTES
	e.GET("/", baseRouteHandler)
	e.GET("/transliterate", controllers.TransliterateController)
	e.POST("/upload", controllers.UploadController)
	e.GET("/upload", uploadRouteHandler)
	e.GET("/signup", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "SIGNUP")
	})
	e.Static("/tmp", "tmp")

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
