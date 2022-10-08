package JsonSchema

import (
	"brinch/lib/constants"
	"brinch/lib/utils"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/cobra"
)

type StoredProcedures struct {
}

func (e StoredProcedures) GetQuery() string {
	return constants.ListAllCustomTypes
}

func (e StoredProcedures) QueryHandler(rows pgx.Rows) utils.Schemas {
	sps := make(map[string]utils.StoredProcedure)
	var schemas []utils.JSONSchemaBase
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

	for _, value := range sps {
		schemas = append(schemas, value.ToJsonSchema())
	}

	return schemas
}
