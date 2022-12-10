package bin

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func MustOpenDbConnection(config *Config) *sql.DB {
	db, err := sql.Open(config.Db.Engine, config.Db.Url)
	if err != nil {
		Logger.Fatal(err)
	}

	return db
}
