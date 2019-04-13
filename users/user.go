package users

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

// CreateUser func
func CreateUser(c echo.Context, conn *sql.DB) error {
	first := c.FormValue("firstname")
	last := c.FormValue("lastname")
	email := c.FormValue("email")
	password := c.FormValue("password")
	crypt, err := hashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	u := &User{
		FirstName: first,
		LastName:  last,
		Email:     email,
		Pass:      crypt,
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
