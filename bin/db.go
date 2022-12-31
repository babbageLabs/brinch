package bin

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"               // this is necessary for connection to postgres db
	_ "github.com/microsoft/go-mssqldb" // this is necessary for connection to mssqldb db
)

func MustOpenDBConnection(config *Config) *sql.DB {
	db, err := sql.Open(config.DB.Engine, config.DB.URL)
	if err != nil {
		if config.App.IsTest() {
			Logger.Error(err)
			db, _, _ = sqlmock.New()
		} else {
			Logger.Fatal(err)
		}
	}

	return db
}
