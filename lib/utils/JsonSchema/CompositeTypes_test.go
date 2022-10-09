package JsonSchema

import (
	"brinch/lib/constants"
	"brinch/lib/utils"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetQuery(t *testing.T) {
	composite := CompositeTypes{}
	query := composite.GetQuery()

	if query != constants.ListAllCustomTypes {
		t.Errorf("composite.GetQuery() = %s; want %s", query, constants.ListAllCustomTypes)
	}
}

func TestDbCustomTypeToJson(t *testing.T) {
	customTypes := make(map[string]utils.CustomTypeAttrs)

	attr := utils.CustomTypeAttr{
		AttrName:         "email",
		TypeName:         "login_tt",
		TypeCategory:     utils.CompositeType,
		AttrTypeName:     "text",
		AttrTypeCategory: utils.StringType,
	}
	attr2 := utils.CustomTypeAttr{
		AttrName:         "username",
		TypeName:         "login_tt",
		TypeCategory:     utils.CompositeType,
		AttrTypeName:     "text",
		AttrTypeCategory: utils.StringType,
	}
	customTypes[attr.TypeName] = append([]utils.CustomTypeAttr{}, attr)
	customTypes[attr.TypeName] = append(customTypes[attr.TypeName], attr2)

	schemas := DbCustomTypeToJsonSchema(&customTypes)

	if len(schemas) != 1 {
		t.Fatalf("expected 1 json schema, got %d", len(schemas))
	}

	schema := schemas[0]
	if len(schema.Properties) != 2 {
		t.Fatalf("expected 2 json schema properties, got %d", len(schema.Properties))
	}
}

func TestGetCustomTypes(t *testing.T) {
	t.Parallel()

	userCategory := CompositeTypes{}

	rows := pgxpoolmock.NewRows([]string{"attr_name", "type_name", "type_category", "attr_type_name", "attr_type_category"}).
		AddRow("username", "user_create", utils.CompositeType, "varchar", utils.StringType).
		AddRow("email", "user_create", utils.CompositeType, "varchar", utils.StringType).
		AddRow("array", "user_create", utils.CompositeType, "varchar", utils.ArrayType).
		AddRow("isMinor", "user_create", utils.CompositeType, "varchar", utils.BooleanType).
		AddRow("age", "user_create", utils.CompositeType, "varchar", utils.NumericType).
		ToPgxRows()
	schemas := userCategory.QueryHandler(rows)

	str, _ := utils.String.ToString()
	arr, _ := utils.Array.ToString()
	boolean, _ := utils.Boolean.ToString()
	numeric, _ := utils.Number.ToString()

	assert.Equal(t, len(schemas), 1)
	assert.Equal(t, len(schemas[0].Properties), 5)
	assert.Equal(t, schemas[0].Properties["username"].Type, str)
	assert.Equal(t, schemas[0].Properties["email"].Type, str)
	assert.Equal(t, schemas[0].Properties["array"].Type, arr)
	assert.Equal(t, schemas[0].Properties["isMinor"].Type, boolean)
	assert.Equal(t, schemas[0].Properties["age"].Type, numeric)
}
