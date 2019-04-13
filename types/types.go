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
