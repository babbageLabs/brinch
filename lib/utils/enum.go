package utils

import (
	"brinch/lib/constants"
	"github.com/jackc/pgx/v4"
)

type Enums struct {
	enums map[string][]string
}

func (enums *Enums) GetQuery() string {
	return constants.ListEnums
}

func (enums *Enums) QueryHandler(rows pgx.Rows) (bool, error) {
	enums.enums = make(map[string][]string)

	for rows.Next() {
		var enumtype string
		var enumlabel string
		err := rows.Scan(&enumtype, &enumlabel)
		if err != nil {
			return false, err
		}

		_, ok := enums.enums[enumtype]

		if !ok {
			var s []string
			enums.enums[enumtype] = append(s, enumlabel)
		} else {
			enums.enums[enumtype] = append(enums.enums[enumtype], enumlabel)
		}
	}
	return true, nil
}

// ToJsonSchema accepts a map of schema names and the enum types and returns a collection of JsonSchema Objects
func (enums *Enums) ToJsonSchema() (Schemas, error) {
	var schemas []JSONSchemaBase
	for k, v := range enums.enums {
		schemas = append(schemas, JSONSchemaBase{
			Id:          k,
			Description: "",
			SchemaType:  String,
			Enum:        v,
		})
	}

	return schemas, nil
}
