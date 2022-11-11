package utils

import (
	"brinch/lib/constants"
	"github.com/jackc/pgx/v4"
)

type StoredProcedures struct {
	sps map[string]StoredProcedure
}

func (procedures *StoredProcedures) GetQuery() string {
	return constants.ListAllSps
}

func (procedures *StoredProcedures) QueryHandler(rows pgx.Rows, meta *DbMeta) (bool, error) {
	procedures.sps = make(map[string]StoredProcedure)
	count := 0

	for rows.Next() {
		count = count + 1
		p := StoredProcedureParameter{}
		err := rows.Scan(&p.RoutineSchema, &p.SpecificName, &p.RoutineName, &p.ParameterName,
			&p.ParameterMode, &p.DataType, &p.UdtName, &p.ParameterDefault)
		if err != nil {
			return false, err
		}
		p.Source = meta.engine
		sp, ok := procedures.sps[p.RoutineName]
		if ok {
			sp.Parameters = append(sp.Parameters, p)
			sp.Source = meta.engine
			procedures.sps[p.RoutineName] = sp
		} else {
			sp = StoredProcedure{
				Name:       p.RoutineName,
				Parameters: append([]StoredProcedureParameter{}, p),
				Source:     meta.engine,
			}
			procedures.sps[p.RoutineName] = sp
		}
	}

	return true, nil
}

func (procedures *StoredProcedures) ToJsonSchema() (Schemas, error) {
	var schemas []JSONSchemaBase

	for _, value := range procedures.sps {
		schemas = append(schemas, value.ToJsonSchema())
	}

	return schemas, nil
}
