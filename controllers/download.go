package controllers

import (
	"fmt"
	"path/filepath"

	"github.com/labstack/echo"
)

// DownloadFile controller
func DownloadFile(c echo.Context) error {
	user, dir, file := c.Param("user"), c.Param("dir"), c.Param("file")
	fmt.Println(user)
	pathToFile := filepath.Join("tmp", dir, file)
	return c.File(pathToFile)
}
