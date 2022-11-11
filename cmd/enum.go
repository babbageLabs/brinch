// Package cmd /*
package cmd

import (
	"brinch/lib/utils"
	"brinch/lib/utils/JsonSchema"
	"github.com/spf13/cobra"
)

// enumCmd represents the enum command
var enumCmd = &cobra.Command{
	Use:   "enum",
	Short: "Export all enums defined in the db to jsonschema",
	Long:  `discover all enum types in the target database and generate json schema representations`,
	Run: func(cmd *cobra.Command, args []string) {
		var enums utils.Enums
		JsonSchema.Export(&enums)
	},
}

func init() {
	scanCmd.AddCommand(enumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enumCmd.PersistentFlags().ToString("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
