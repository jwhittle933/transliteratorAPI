package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// CreateUser func
func CreateUser(c echo.Context, conn *sql.DB) error {
	u := &User{}
	email := c.FormValue("email")
	password := c.FormValue("password")
	fmt.Println(email)
	fmt.Println(password)

	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println(u)
	return c.JSON(http.StatusOK, http.StatusOK)
}

// GetUser func
func GetUser(c echo.Context, conn *sql.DB) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	return nil
}

// UpdateUser func
func UpdateUser(c echo.Context, conn *sql.DB) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return nil
}

// DeleteUser func
func DeleteUser(c echo.Context, conn *sql.DB) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	return c.NoContent(http.StatusNoContent)
}
