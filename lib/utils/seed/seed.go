package seed

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/fs"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ScanDir(path string, db *sql.DB) {
	fileMatchPattern := viper.GetString("db.fileMatchPattern")
	//fileKey := fmt.Sprintf("brinch.%s.files", appName)

	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		cobra.CheckErr(err)

		if !info.IsDir() {
			match, _ := regexp.MatchString(fileMatchPattern, info.Name())
			if match {
				priority, err := strconv.ParseFloat(strings.Split(info.Name(), ".")[0], 64)
				cobra.CheckErr(err)

				fmt.Printf("%s , %f \n", path, priority)
				_, err = FileSeed(path, db)
				if err != nil {
					fmt.Printf("%s : %s \n", "error seeding file", path)
					cobra.CheckErr(err)
				}
			}
		}
		return nil
	})
	cobra.CheckErr(err)
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
