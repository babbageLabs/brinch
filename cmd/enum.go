/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"brinch/lib/constants"
	"brinch/lib/utils"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
)

// enumCmd represents the enum command
var enumCmd = &cobra.Command{
	Use:   "enum",
	Short: "Discover all enum types in db",
	Long:  `discover all enum types in the target database and generate json schema representations`,
	Run: func(cmd *cobra.Command, args []string) {
		query := constants.ListEnums
		utils.QueryDB(&query, handleQueryEnums)
	},
}

func init() {
	scanCmd.AddCommand(enumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleQueryEnums(rows pgx.Rows) {
	var ret = make(map[string][]string)

	for rows.Next() {
		var enumtype string
		var enumlabel string
		err := rows.Scan(&enumtype, &enumlabel)
		if err != nil {
			logrus.Error(err)
		}

		_, ok := ret[enumtype]

		if ok != true {
			var s []string
			ret[enumtype] = append(s, enumlabel)
		} else {
			ret[enumtype] = append(ret[enumtype], enumlabel)
		}
	}
	fmt.Printf("No of User Defined enums found: %d \n", len(ret))
	dbEnumToJsonSchema(&ret)
}

// dbEnumToJsonSchema accepts a map of schema names and the enum types and returns a collection of JsonSchema Objects
func dbEnumToJsonSchema(values *map[string][]string) {
	schema := viper.GetString("jsonSchema.schema")
	path := viper.GetString("jsonSchema.targetPath")

	if path == "" {
		path, err := os.Getwd()
		cobra.CheckErr(err)

		path = filepath.Join(path, "schema", "json")
	}

	for k, v := range *values {
		id := filepath.Join(path, k+".schema"+".json")

		sc := utils.JSONSchemaBase{
			Schema:      schema,
			Id:          id,
			Name:        k + ".schema" + ".json",
			Title:       cases.Title(language.English, cases.Compact).String(k),
			Description: "",
			SchemaType:  utils.String,
			Enum:        v,
		}

		sc.ToFile()
	}
}
