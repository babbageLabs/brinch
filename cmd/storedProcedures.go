/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"brinch/lib/constants"
	"brinch/lib/utils"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/cobra"
)

var CustomTypes = make(map[string]utils.CustomTypes)

// storedProceduresCmd represents the storedProcedures command
var storedProceduresCmd = &cobra.Command{
	Use:   "storedProcedures",
	Short: "Scan all stored procedures",
	Long:  `A command to discover all composite types, domains and stored procedures`,
	Run: func(cmd *cobra.Command, args []string) {
		query := constants.ListAllCustomTypes
		utils.QueryDB(&query, getCustomTypes)

		query = constants.ListAllSps
		utils.QueryDB(&query, getSps)
	},
}

func init() {
	scanCmd.AddCommand(storedProceduresCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storedProceduresCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storedProceduresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getSps(rows pgx.Rows) {
	sps := make(map[string]utils.StoredProcedure)
	count := 0

	for rows.Next() {
		count = count + 1
		p := utils.StoredProcedureParameter{}

		err := rows.Scan(&p.RoutineSchema, &p.SpecificName, &p.RoutineName, &p.ParameterName,
			&p.ParameterMode, &p.DataType, &p.UdtName, &p.ParameterDefault)
		cobra.CheckErr(err)

		sp, ok := sps[p.RoutineName]
		if ok == true {
			sp.Parameters = append(sp.Parameters, p)
			sps[p.RoutineName] = sp
		} else {
			sp = utils.StoredProcedure{
				Name:       p.RoutineName,
				Parameters: append([]utils.StoredProcedureParameter{}, p),
			}
			sps[p.RoutineName] = sp
		}
	}
	fmt.Printf("No of Stored Proceedures found: %d \n", len(sps))
	scanSPParams(sps)
}

func getCustomTypes(rows pgx.Rows) {
	count := 0

	for rows.Next() {
		count = count + 1
		attr := utils.CustomTypeAttr{}

		err := rows.Scan(&attr.AttrName, &attr.TypeName, &attr.TypeCategory, &attr.AttrTypeName, &attr.AttrTypeCategory)

		cobra.CheckErr(err)

		typ := utils.CustomTypes{
			Name:         attr.TypeName,
			TypeCategory: attr.TypeCategory,
		}

		CustomTypes[typ.Name] = typ
	}
	fmt.Printf("No of User Defined types found: %d \n", len(CustomTypes))
}

func scanSPParams(sps map[string]utils.StoredProcedure) {

	for _, value := range sps {
		schema := value.ToJsonSchema()
		schema.ResolvePropertyTypes(CustomTypes)
		schema.ToFile()
	}
}
