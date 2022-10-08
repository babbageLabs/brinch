package JsonSchema

import (
	"brinch/lib/constants"
	"brinch/lib/utils"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Enums struct {
}

func (e Enums) GetQuery() string {
	return constants.ListEnums
}

func (e Enums) QueryHandler(rows pgx.Rows) utils.Schemas {
	var ret = make(map[string][]string)

	for rows.Next() {
		var enumtype string
		var enumlabel string
		err := rows.Scan(&enumtype, &enumlabel)
		if err != nil {
			logrus.Error(err)
		}

		_, ok := ret[enumtype]

		if ok != true {
			var s []string
			ret[enumtype] = append(s, enumlabel)
		} else {
			ret[enumtype] = append(ret[enumtype], enumlabel)
		}
	}
	fmt.Printf("No of User Defined enums found: %d \n", len(ret))
	return dbEnumToJsonSchema(&ret)
}

// dbEnumToJsonSchema accepts a map of schema names and the enum types and returns a collection of JsonSchema Objects
func dbEnumToJsonSchema(values *map[string][]string) utils.Schemas {
	var enums []utils.JSONSchemaBase
	for k, v := range *values {
		enums = append(enums, utils.JSONSchemaBase{
			Id:          k,
			Description: "",
			SchemaType:  utils.String,
			Enum:        v,
		})
	}

	return enums
}
