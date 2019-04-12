package controllers

import (
	"path/filepath"

	"github.com/labstack/echo"
)

// DownloadFile controller
func DownloadFile(c echo.Context) error {
	user, dir, file := c.Param("user"), c.Param("dir"), c.Param("file")
	pathToFile := filepath.Join("tmp", "users", user, dir, file)
	return c.File(pathToFile)
}
