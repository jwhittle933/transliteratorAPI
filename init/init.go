package start

import (
	"database/sql"

	"github.com/labstack/echo"
)

// Init func for instantiating
func Init() (*echo.Echo, *sql.DB, error) {
	e := echo.New()

	conn, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return nil, nil, err
	}

	return e, conn, nil
}
