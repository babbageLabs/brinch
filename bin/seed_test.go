package bin

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeed_ScanSuccess(t *testing.T) {
	seed := &Seed{
		Path:             "../testdata",
		FileMatchPattern: "^.*\\.(sql)$",
	}

	files, err := seed.Scan()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(files))
}

func TestSeed_ScanNonExistentPath(t *testing.T) {
	seed := &Seed{
		Path:             "../testdatas",
		FileMatchPattern: "^.*\\.(sql)$",
	}

	_, err := seed.Scan()
	assert.Error(t, err)
}

func TestSeed_ScanInvalidRegex(t *testing.T) {
	seed := &Seed{
		Path:             "../testdata",
		FileMatchPattern: "`[ ]\\K(?<!\\d )(?=(?: ?\\d){8})(?!(?: ?\\d){9})\\d[ \\d]+\\d`",
	}

	_, err := seed.Scan()
	assert.Error(t, err)
}

func TestSeed_SeedSequentialSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			Logger.Fatal(err)
		}
	}(db)

	// Initialize SqlFile
	s := NewSQLFile()
	e := s.LoadFile("../testdata/expected.sql")
	assert.NoError(t, e)
	// before we actually execute our api function, we need to expect required DB actions
	var lastInsertID, affected int64
	rows := sqlmock.NewResult(lastInsertID, affected)
	mock.ExpectBegin()
	mock.ExpectExec(s.queries[0]).WillReturnResult(rows)
	mock.ExpectCommit()
	mock.ExpectClose()

	seed := &Seed{
		Path:             "../testdata",
		FileMatchPattern: "^expected\\.sql$",
		DB:               db,
		Mode:             Sequential,
	}

	_, e = seed.Seed()
	assert.NoError(t, e)
}

func TestSeed_SeedBatchedSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			Logger.Fatal(err)
		}
	}(db)

	// Initialize SqlFile
	s := NewSQLFile()
	e := s.LoadFile("../testdata/expected.sql")
	assert.NoError(t, e)
	// before we actually execute our api function, we need to expect required DB actions
	//var lastInsertID, affected int64
	//rows := sqlmock.NewResult(lastInsertID, affected)
	//mock.ExpectBegin()
	//mock.ExpectExec(s.queries[0]).WillReturnResult(rows)
	//mock.ExpectCommit()
	mock.ExpectClose()

	seed := &Seed{
		Path:             "../testdata",
		FileMatchPattern: "^expected\\.sql$",
		DB:               db,
		Mode:             Batched,
	}

	_, e = seed.Seed()
	assert.NoError(t, e)
}
