package JsonSchema

import (
	"brinch/lib/constants"
	"brinch/lib/utils"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/cobra"
)

type CompositeType struct {
}

func (e CompositeType) GetQuery() string {
	return constants.ListAllCustomTypes
}

func (e CompositeType) QueryHandler(rows pgx.Rows) utils.Schemas {
	count := 0
	customTypes := make(map[string]utils.CustomTypeAttrs)

	for rows.Next() {
		count = count + 1
		attr := utils.CustomTypeAttr{}

		err := rows.Scan(&attr.AttrName, &attr.TypeName, &attr.TypeCategory, &attr.AttrTypeName, &attr.AttrTypeCategory)
		cobra.CheckErr(err)

		tt, ok := customTypes[attr.TypeName]

		if ok {
			customTypes[attr.TypeName] = append(tt, attr)
		} else {
			customTypes[attr.TypeName] = append([]utils.CustomTypeAttr{}, attr)
		}
	}
	fmt.Printf("No of User Defined types found: %d \n", len(customTypes))

	return dbCustomTypeToJsonSchema(&customTypes)
}

func dbCustomTypeToJsonSchema(values *map[string]utils.CustomTypeAttrs) utils.Schemas {
	var customTypes utils.Schemas
	for typId, v := range *values {
		fmt.Printf("%s\n", typId)
		customTypes = append(customTypes, utils.JSONSchemaBase{
			Id:          typId,
			Description: "A composite type",
			SchemaType:  utils.Object,
			Properties:  v.ToJsonSchemaProperties(),
		})
	}

	return customTypes
}
