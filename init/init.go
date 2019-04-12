package start

import (
	"database/sql"

	"github.com/labstack/echo"
)

// Init func for instantiating
func Init() (*echo.Echo, *sql.DB, error) {
	e := echo.New()

	/*
	 * sql.Open() doesn't directly open a conection
	 * and won't return error if the server isn't
	 * available or the conn data isn't correct.
	 *
	 * Thus, sql.Ping() is used to check for err
	 *
	 * MySQL driver imported in main
	 */
	conn, _ := sql.Open("mysql", "user:password@/dbname")
	err := conn.Ping()
	if err != nil {
		return nil, nil, err
	}

	return e, conn, nil
}
