package utils

import (
	"brinch/lib/constants"
	"brinch/lib/utils/databases"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuery(t *testing.T) {
	t.Parallel()
	composite := CompositeTypes{}
	query := composite.GetQuery()

	if query != constants.ListAllCustomTypes {
		t.Errorf("composite.GetQuery() = %s; want %s", query, constants.ListAllCustomTypes)
	}
}

func TestGetCustomTypes(t *testing.T) {
	t.Parallel()

	userCreate := CompositeTypes{}
	dbMeta := databases.DbMeta{}
	typeName := "user_create"

	rows := pgxpoolmock.NewRows([]string{"attr_name", "type_name", "type_category", "attr_type_name", "attr_type_category"}).
		AddRow("username", typeName, databases.CompositeType, "varchar", databases.StringType).
		AddRow("email", typeName, databases.CompositeType, "varchar", databases.StringType).
		AddRow("array", typeName, databases.CompositeType, "varchar", databases.ArrayType).
		AddRow("isMinor", typeName, databases.CompositeType, "varchar", databases.BooleanType).
		AddRow("age", typeName, databases.CompositeType, "varchar", databases.NumericType).
		ToPgxRows()
	_, err := userCreate.QueryHandler(rows, &dbMeta)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(userCreate.types), 1)

	properties := userCreate.types["user_create"].ToJsonSchemaProperties()

	//str, _ := JsonSchema2.String.ToString()
	//arr, _ := JsonSchema2.Array.ToString()
	//boolean, _ := JsonSchema2.Boolean.ToString()
	//numeric, _ := JsonSchema2.Number.ToString()
	assert.Equal(t, len(properties), 5)
	//assert.Equal(t, properties["username"].Type, str)
	//assert.Equal(t, properties["email"].Type, str)
	//assert.Equal(t, properties["array"].Type, arr)
	//assert.Equal(t, properties["isMinor"].Type, boolean)
	//assert.Equal(t, properties["age"].Type, numeric)

	schemas, err := userCreate.ToJsonSchema()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(schemas), 1)
}
