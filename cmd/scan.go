/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"brinch/lib/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/fs"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Locate sql files and generate db initialization config",
	Long: `Locate all sql files in the current directory and generate a
database initialization configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		schemas := viper.GetStringSlice("db.schema")

		for _, element := range schemas {
			logrus.Info("Scanning path ", element)
			scanDir(element)
			logrus.Info("Scanning completed")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func scanDir(path string) {
	conn, ctx := utils.CreateConnection()
	appName := viper.GetString("app.name")
	fileKey := fmt.Sprintf("brinch.%s.files", appName)

	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() {
			match, _ := regexp.MatchString("\\d+\\.\\w+\\.\\w+\\.\\w+\\.(sql|yaml|yml)", info.Name())
			if match {
				priority, err := strconv.ParseFloat(strings.Split(info.Name(), ".")[0], 64)
				if err != nil {
					logrus.Error(err)
					panic(err)
				}

				err = conn.ZAdd(ctx, fileKey, &redis.Z{
					Score:  priority,
					Member: path,
				}).Err()
				if err != nil {
					logrus.Error(err)
					panic(err)
				}

				fmt.Printf("%s\n", path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
