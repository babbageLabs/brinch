package seed

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
)

func TestFileLoad__success(t *testing.T) {
	t.Helper()
	path := "../../../testdata/include_comments.sql"
	// Get a database handler
	db, mock := newMockDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s %v \n", "error", err)
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

func TestFileScan__success(t *testing.T) {
	// Get a database handler
	db, _ := newMockDB(t)
	fileMatchPattern := "(?s).*"
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s %v \n", "error", err)
		}
	}(db)
	path := "../../../testdata"

	_, err := ScanDir(&path, &fileMatchPattern, db)
	if err != nil {
		fmt.Printf("%s %v \n", "error", err)
	}
}
