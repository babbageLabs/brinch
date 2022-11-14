package utils

import (
	"brinch/lib/constants"
	JsonSchema2 "brinch/lib/utils/JsonSchema"
	"brinch/lib/utils/databases"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/cobra"
)

type CompositeTypes struct {
	types map[string]CustomTypeAttrs
}

func (composites *CompositeTypes) QueryHandler(rows pgx.Rows, meta *databases.DbMeta) (bool, error) {
	count := 0
	composites.types = make(map[string]CustomTypeAttrs)

	for rows.Next() {
		count = count + 1
		attr := CustomTypeAttr{
			Source: meta.SourceType,
		}

		err := rows.Scan(&attr.AttrName, &attr.TypeName, &attr.TypeCategory, &attr.AttrTypeName, &attr.AttrTypeCategory)
		cobra.CheckErr(err)

		tt, ok := composites.types[attr.TypeName]
		if ok {
			composites.types[attr.TypeName] = append(tt, attr)
		} else {
			composites.types[attr.TypeName] = append([]CustomTypeAttr{}, attr)
		}
	}

	return true, nil
}

func (composites *CompositeTypes) GetQuery() string {
	return constants.ListAllCustomTypes
}

func (composites *CompositeTypes) ToJsonSchema() (JsonSchema2.Schemas, error) {
	var customTypes JsonSchema2.Schemas
	for typId, v := range composites.types {
		customTypes = append(customTypes, JsonSchema2.JSONSchemaBase{
			Id:          typId,
			Description: "A composite type",
			SchemaType:  JsonSchema2.Object,
			Properties:  v.ToJsonSchemaProperties(),
		})
	}

	return customTypes, nil
}
