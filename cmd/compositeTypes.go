// Package cmd /*
package cmd

import (
	"brinch/lib/utils"
	"brinch/lib/utils/JsonSchema"
	"github.com/spf13/cobra"
)

// compositeTypesCmd represents the compositeTypes command
var compositeTypesCmd = &cobra.Command{
	Use:   "compositeTypes",
	Short: "Export all composite types defined in the db to jsonschema",
	Long: ` Get a list of all custom types defined in the database
	and generate the representative json schema representation
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var composites utils.CompositeTypes
		JsonSchema.Export(&composites)
	},
}

func init() {
	scanCmd.AddCommand(compositeTypesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compositeTypesCmd.PersistentFlags().ToString("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compositeTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
