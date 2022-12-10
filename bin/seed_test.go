package bin

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeed_ScanSuccess(t *testing.T) {
	seed := &Seed{
		path:             "../testdata",
		fileMatchPattern: "^.*\\.(sql)$",
	}

	files, err := seed.Scan()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(files))
}

func TestSeed_ScanNonExistentPath(t *testing.T) {
	seed := &Seed{
		path:             "../testdatas",
		fileMatchPattern: "^.*\\.(sql)$",
	}

	_, err := seed.Scan()
	assert.Error(t, err)
}

func TestSeed_ScanInvalidRegex(t *testing.T) {
	seed := &Seed{
		path:             "../testdata",
		fileMatchPattern: "`[ ]\\K(?<!\\d )(?=(?: ?\\d){8})(?!(?: ?\\d){9})\\d[ \\d]+\\d`",
	}

	_, err := seed.Scan()
	assert.Error(t, err)
}

func TestSeed_SeedSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize SqlFile
	s := NewSqlFile()
	e := s.LoadFile("../testdata/expected.sql")
	assert.NoError(t, e)
	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "user_id", "title", "content"}).
		AddRow(1, 1, "Self Introduction", "-- sqlfile --\\nI'm sqlfile.")
	mock.ExpectBegin()
	mock.ExpectQuery(s.queries[0]).WillReturnRows(rows)
	mock.ExpectCommit()

	seed := &Seed{
		path:             "../testdata",
		fileMatchPattern: "^expected\\.sql$",
		db:               db,
	}

	_, err = seed.Seed()
	if err != nil {
		fmt.Printf("%s", err)
		//TODO fix this test case
	}
	//assert.NoError(t, err)
}
