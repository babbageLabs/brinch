// Package cmd /*
package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// serveSchemaCmd represents the serveShema command
var serveSchemaCmd = &cobra.Command{
	Use:   "serveSchema",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running schema server...")
		fileServer()
	},
}

func init() {
	rootCmd.AddCommand(serveSchemaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveSchemaCmd.PersistentFlags().ToString("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveSchemaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fileServer() {
	port := viper.GetString("jsonSchema.server.port")
	directory := viper.GetString("jsonSchema.targetPath")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
