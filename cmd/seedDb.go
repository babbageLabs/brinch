// Package cmd /*
package cmd

import (
	"brinch/lib/utils/seed"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// seedDbCmd represents the seedDb command
var seedDbCmd = &cobra.Command{
	Use:   "seedDb",
	Short: "Discover",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		schemas := viper.GetStringSlice("db.schemas")
		engine := viper.GetString("db.config.engine")
		url := viper.GetString("db.config.url")
		fileMatchPattern := viper.GetString("db.fileMatchPattern")

		// Get a database handler
		db, err := sql.Open(engine, url)
		cobra.CheckErr(err)

		fmt.Printf("%s \n", schemas)
		for _, element := range schemas {
			fmt.Printf("Scanning files in path %s \n", element)
			_, err := seed.ScanDir(&element, &fileMatchPattern, db)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedDbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedDbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedDbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
