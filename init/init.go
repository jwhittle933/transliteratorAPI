package start

import (
	"database/sql"

	"github.com/jwhittle933/transliteratorAPI/types"
	"github.com/labstack/echo"
)

// Init func for instantiating
func Init() (*types.AppMeta, error) {
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
		return nil, err
	}

	app := &types.AppMeta{
		Echo: e,
		DB:   conn,
	}

	return app, nil
}
