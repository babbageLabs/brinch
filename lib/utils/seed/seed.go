package seed

import (
	"database/sql"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ScanDir(path *string, fileMatchPattern *string, db *sql.DB) (string, error) {
	err := filepath.WalkDir(*path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			match, _ := regexp.MatchString(*fileMatchPattern, info.Name())
			if match {
				priority, err := strconv.ParseFloat(strings.Split(info.Name(), ".")[0], 64)
				if err != nil {
					priority = -1
				}

				fmt.Printf("%s , %f \n", path, priority)
				_, err = FileSeed(path, db)
				if err != nil {
					fmt.Printf("%s : %s \n", "error seeding file", path)
					return nil
				}
			}
		}
		return nil
	})
	return "", err
}

func FileSeed(path string, db *sql.DB) (res []sql.Result, err error) {
	// Initialize SqlFile
	s := NewSqlFile()
	// Load input file and store queries written in the file
	e := s.LoadFile(path)
	if e != nil {
		return nil, err
	}
	// Execute the stored queries
	// transaction is used to execute queries in Exec()
	r, e := s.Exec(db)

	return r, e
}
