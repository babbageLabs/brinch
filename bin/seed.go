package bin

import (
	"database/sql"
	"io/fs"
	"path/filepath"
	"regexp"
)

type Seed struct {
	Path             string
	FileMatchPattern string
	DB               *sql.DB
}

// Scan traverse a path and discover or paths that match a specific pattern in lexical order
func (seed *Seed) Scan() ([]string, error) {
	var files []string

	err := filepath.WalkDir(seed.Path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			match, err := regexp.MatchString(seed.FileMatchPattern, info.Name())
			if err != nil {
				return err
			}
			if match {
				Logger.Info("Collecting file to seed to seed ", info.Name())
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
		Logger.Info("Preparing to seed ", path)
		_, err := seed.SeedFile(path)
		if err != nil {
			Logger.Info("Seeding failed for ", path, " with error ", err)
			return false, err
		}
		Logger.Info("Seeding completed for ", path)
	}

	return true, nil
}

func (seed *Seed) SeedFile(path string) (res []sql.Result, err error) {
	// Initialize SqlFile
	s := NewSqlFile()
	// Load input file and store queries written in the file
	e := s.LoadFile(path)
	if e != nil {
		return nil, err
	}
	// Execute the stored queries
	// transaction is used to execute queries in Exec()
	r, e := s.Exec(seed.DB)

	return r, e
}
