package transports

import (
	"database/sql"
	"fmt"
	"github.com/babbageLabs/brinch/bin/core/types"
)

type DBTransport struct {
	DB *sql.DB
}

func (db *DBTransport) Connect() (bool, error) {
	if db.DB != nil {
		return true, nil
	}
	return false, fmt.Errorf("database connection is not provided")
}

func (db *DBTransport) Close() (bool, error) {
	_, err := db.Connect()
	if err != nil {
		return false, err
	}

	err = db.DB.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *DBTransport) Exec(subject string, msg []byte, meta *types.MetaData) (*types.Response, error) {
	// TODO implement me
	return nil, nil
}
