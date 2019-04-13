package types

import (
	"database/sql"
	"mime/multipart"

	"github.com/labstack/echo"
)

// AppMeta struct for app level data exchange
type AppMeta struct {
	Echo *echo.Echo
	DB   *sql.DB
}

// User struct for registering users
type User struct {
	ID        int    `json:"id" form:"id" query:"id"`
	FirstName string `json:"firstname" form:"firstname" query:"firstname"`
	LastName  string `json:"lastname" form:"lastname" query:"lastname"`
	Email     string `json:"email" form:"email" query:"email"`
	Pass      string `json:"password" form:"password" query:"password"` //encrypt
}

// Resp struct for basic response schema.
type Resp struct {
	Code    int64
	Message string
}

// File struct
type File struct {
	File        multipart.File
	Mime        string
	Size        int64
	ByteContent []byte
	TextContext string
}
