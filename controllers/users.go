package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

// CreateUser func
func CreateUser(c echo.Context) error {
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
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	return nil
}

// UpdateUser func
func UpdateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return nil
}

// DeleteUser func
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	return c.NoContent(http.StatusNoContent)
}

// LoginUser user
func LoginUser() {
	//
}

// LogoutUser user
func LogoutUser() {
	//
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
