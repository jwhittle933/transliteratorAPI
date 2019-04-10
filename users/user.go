package users

import (
	// "github.com/jwhittle933/transliteratorAPI/types"
	"fmt"
	"net/http"

	"../types"
	"github.com/labstack/echo"
)

// CreateUser func
func CreateUser(c echo.Context) error {
	u := &types.User{}

	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println(u)
	return c.JSON(http.StatusOK, &types.Resp{
		Code:    http.StatusOK,
		Message: "User successfully created.",
	})
}

// GetUser func
func GetUser(c echo.Context) error {
	return nil
}

// UpdateUser func
func UpdateUser(c echo.Context) error {
	return nil
}

// DeleteUser func
func DeleteUser(c echo.Context) error {
	return nil
}
