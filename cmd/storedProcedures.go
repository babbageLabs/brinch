// Package cmd /*
package cmd

import (
	"brinch/lib/utils/JsonSchema"
	"github.com/spf13/cobra"
)

// storedProceduresCmd represents the storedProcedures command
var storedProceduresCmd = &cobra.Command{
	Use:   "storedProcedures",
	Short: "Export all stored procedures defined in the db to jsonschema",
	Long:  `A command to discover all composite types, domains and stored procedures`,
	Run: func(cmd *cobra.Command, args []string) {
		JsonSchema.Export(JsonSchema.StoredProcedures{})
	},
}

func init() {
	scanCmd.AddCommand(storedProceduresCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storedProceduresCmd.PersistentFlags().ToString("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storedProceduresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
