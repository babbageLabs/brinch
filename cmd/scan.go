// Package cmd /*
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			//seed.ScanDir(element)
			logrus.Info("Scanning completed")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().ToString("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//func scanDir(path string) {
//	conn, ctx := utils.CreateConnection()
//	appName := viper.GetString("app.name")
//	fileKey := fmt.Sprintf("brinch.%s.files", appName)
//
//	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
//		cobra.CheckErr(err)
//
//		if !info.IsDir() {
//			match, _ := regexp.MatchString("\\d+\\.\\w+\\.\\w+\\.\\w+\\.(sql|yaml|yml)", info.Name())
//			if match {
//				priority, err := strconv.ParseFloat(strings.Split(info.Name(), ".")[0], 64)
//				cobra.CheckErr(err)
//
//				err = conn.ZAdd(ctx, fileKey, &redis.Z{
//					Score:  priority,
//					Member: path,
//				}).Err()
//				cobra.CheckErr(err)
//
//				fmt.Printf("%s\n", path)
//			}
//		}
//		return nil
//	})
//	cobra.CheckErr(err)
//}
