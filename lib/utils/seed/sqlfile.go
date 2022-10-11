// Package seed provides a way to execute sql file easily
//
// For more usage see https://github.com/tanimutomo/sqlfile

package seed

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// SqlFile represents a queries holder
type SqlFile struct {
	files   []string
	queries []string
}

// NewSqlFile create new SqlFile object
func NewSqlFile() *SqlFile {
	return &SqlFile{}
}

// LoadFile add and load queries from input file
func (s *SqlFile) LoadFile(file string) error {
	queries, err := load(file)
	if err != nil {
		return err
	}

	s.files = append(s.files, file)
	s.queries = append(s.queries, queries...)

	return nil
}

// LoadFiles add and load queries from multiple input files
func (s *SqlFile) LoadFiles(files ...string) error {
	for _, file := range files {
		if err := s.LoadFile(file); err != nil {
			return err
		}
	}
	return nil
}

// Directory add and load queries from *.sql files in specified directory
func (s *SqlFile) Directory(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if name[len(name)-3:] != "sql" {
			continue
		}

		if err := s.LoadFile(dir + "/" + name); err != nil {
			return err
		}
	}

	return nil
}

// Exec execute SQL statements written int the specified sql file
func (s *SqlFile) Exec(db *sql.DB) (res []sql.Result, err error) {
	tx, err := db.Begin()
	if err != nil {
		return res, err
	}
	defer saveTx(tx, &err)

	var rs []sql.Result
	for _, q := range s.queries {
		r, err := tx.Exec(q)
		if err != nil {
			return res, fmt.Errorf(err.Error() + " : when executing > " + q)
		}
		rs = append(rs, r)
	}

	return rs, err
}

// Load sql file from path, and return SqlFile pointer
func load(path string) (qs []string, err error) {
	ls, err := readFileContents(path)
	if err != nil {
		return qs, err
	}
	return ls, nil
}

func readFileContents(path string) (ls []string, err error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return ls, err
	}
	var temp []string

	//ls = strings.Split(string(f), "\n")
	return append(temp, string(f)), nil
}

func saveTx(tx *sql.Tx, err *error) {
	if p := recover(); p != nil {
		err := tx.Rollback()
		cobra.CheckErr(err)
		panic(p)
	} else if *err != nil {
		err := tx.Rollback()
		cobra.CheckErr(err)
	} else {
		e := tx.Commit()
		err = &e
	}
}
