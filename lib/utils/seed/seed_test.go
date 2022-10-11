package seed

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
)

func TestFileLoad(t *testing.T) {
	t.Helper()
	path := "../../../testdata/include_comments.sql"
	// Get a database handler
	db, mock := newMockDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	mock.ExpectBegin()

	query, err := load(path)
	if err != nil {
		t.Errorf("%s %v \n", "error reading test data", err)
	}

	mock.ExpectExec(regexp.QuoteMeta(query[0])).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	_, err = FileSeed(path, db)
	if err != nil {
		t.Errorf("%s %v \n", "error ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
