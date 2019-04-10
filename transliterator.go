package main

import (
	"net/http"

	// "github.com/jwhittle933/transliteratorAPI/controllers"
	"./controllers"
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
	e.Static("/tmp", "tmp")

	// auth := e.Group("/auth")
	// auth.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "Joe" && password == "password" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	// START
	e.Logger.Fatal(e.Start(":3000"))
}

func baseRouteHandler(c echo.Context) error {
	resp := &types.Resp{
		Code:    200,
		Message: "Transliterator API",
	}
	return c.JSON(http.StatusOK, resp)
}

func uploadRouteHandler(c echo.Context) error {
	resp := &types.Resp{
		Code:    200,
		Message: "Upload a file",
	}
	return c.JSON(http.StatusOK, resp)
}
