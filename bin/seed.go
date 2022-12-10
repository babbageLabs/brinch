package bin

import (
	"database/sql"
	"io/fs"
	"path/filepath"
	"regexp"
)

type Seed struct {
	path             string
	fileMatchPattern string
	db               *sql.DB
}

// Scan traverse a path and discover or paths that match a specific pattern in lexical order
func (seed *Seed) Scan() ([]string, error) {
	var files []string

	err := filepath.WalkDir(seed.path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			match, err := regexp.MatchString(seed.fileMatchPattern, info.Name())
			if err != nil {
				return err
			}
			if match {
				files = append(files, path)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Seed execute the discovered paths against the DB
func (seed *Seed) Seed() (bool, error) {
	paths, err := seed.Scan()
	if err != nil {
		return false, err
	}

	for _, path := range paths {
		_, err := seed.SeedFile(&path)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (seed *Seed) SeedFile(path *string) (res []sql.Result, err error) {
	// Initialize SqlFile
	s := NewSqlFile()
	// Load input file and store queries written in the file
	e := s.LoadFile(*path)
	if e != nil {
		return nil, err
	}
	// Execute the stored queries
	// transaction is used to execute queries in Exec()
	r, e := s.Exec(seed.db)

	return r, e
}
