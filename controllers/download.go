package controllers

import (
	"path/filepath"

	"github.com/labstack/echo"
)

// DownloadFile controller
func DownloadFile(c echo.Context) error {
	/*
	 * The exact method of storing user files is uncertain
	 *
	 * Perhaps the files will be stored remotely (S3)
	 */
	user, dir, file := c.Param("user"), c.Param("dir"), c.Param("file")
	pathToFile := filepath.Join("tmp", "users", user, dir, file)
	return c.File(pathToFile)
}
