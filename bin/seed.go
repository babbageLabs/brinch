package bin

import (
	"database/sql"
	"io/fs"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type SeedMode string

const (
	Sequential SeedMode = "sequential"
	Batched    SeedMode = "batched"
)

type Seed struct {
	Path             string
	FileMatchPattern string
	DB               *sql.DB
	Mode             SeedMode
}

type SeedPath struct {
	path     string
	fileName string
	weight   int64
}

type SeedPaths []SeedPath

const DefaultFileWeight = 1000

// Scan traverse a path and discover or paths that match a specific pattern in lexical order
func (seed *Seed) Scan() (SeedPaths, error) {
	var files []SeedPath

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
				// Logger.Info("Collecting file to seed to seed ", info.Name())
				files = append(files, SeedPath{
					path:     path,
					fileName: info.Name(),
					weight:   seed.GetFileWeight(info.Name()),
				})
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

	switch seed.Mode {
	case Batched:
		grouped := paths.GroupByPriority()
		for _, value := range grouped {
			seed.ExecBatch(value)
		}
	case Sequential:
		paths = seed.OrderFiles(paths)
		for _, path := range paths {
			Logger.Info("Preparing to seed ", path.path)
			_, err := seed.ExecuteFile(path.path)
			if err != nil {
				Logger.Error("Seeding failed for ", path, " with error ", err)
				return false, err
			}
			Logger.Info("Seeding completed for ", path)
		}
	}
	return true, nil
}

func (seed *Seed) ExecuteFile(path string) (res []sql.Result, err error) {
	// Initialize SqlFile
	s := NewSQLFile()
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

func (seed *Seed) OrderFiles(paths SeedPaths) SeedPaths {
	sort.Slice(paths, func(i int, j int) bool {
		return paths[i].weight < paths[j].weight
	})

	return paths
}

func (seed *Seed) GetFileWeight(fileName string) int64 {
	var weight []string
	for _, v := range fileName {
		if unicode.IsDigit(v) {
			weight = append(weight, string(v))
		} else {
			break
		}
	}

	if len(weight) != 0 {
		weight := strings.Join(weight, "")
		w, err := strconv.Atoi(weight)
		if err == nil {
			return int64(w)
		}
	}
	return DefaultFileWeight
}

func (seed *Seed) ExecBatch(paths SeedPaths) {
	Logger.Debug("Seeding batch", paths[0].weight)
	for _, path := range paths {
		path := path
		go func() {
			Logger.Info("Preparing to seed ", path.path)
			_, err := seed.ExecuteFile(path.path)
			if err != nil {
				Logger.Error("Seeding failed for ", path, " with error ", err)
			} else {
				Logger.Info("Seeding completed for ", path)
			}
		}()
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (paths *SeedPaths) GroupByPriority() map[int64]SeedPaths {
	ret := make(map[int64]SeedPaths)

	for _, path := range *paths {
		_, ok := ret[path.weight]
		if ok {
			ret[path.weight] = append(ret[path.weight], path)
		} else {
			var ws []SeedPath
			ret[path.weight] = append(ws, path)
		}
	}

	return ret
}
