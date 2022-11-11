package utils

import (
	"brinch/lib/constants"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuery(t *testing.T) {
	composite := CompositeTypes{}
	query := composite.GetQuery()

	if query != constants.ListAllCustomTypes {
		t.Errorf("composite.GetQuery() = %s; want %s", query, constants.ListAllCustomTypes)
	}
}

func TestGetCustomTypes(t *testing.T) {
	t.Parallel()

	userCreate := CompositeTypes{}
	typeName := "user_create"

	rows := pgxpoolmock.NewRows([]string{"attr_name", "type_name", "type_category", "attr_type_name", "attr_type_category"}).
		AddRow("username", typeName, CompositeType, "varchar", StringType).
		AddRow("email", typeName, CompositeType, "varchar", StringType).
		AddRow("array", typeName, CompositeType, "varchar", ArrayType).
		AddRow("isMinor", typeName, CompositeType, "varchar", BooleanType).
		AddRow("age", typeName, CompositeType, "varchar", NumericType).
		ToPgxRows()
	_, err := userCreate.QueryHandler(rows)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(userCreate.types), 1)

	properties := userCreate.types["user_create"].ToJsonSchemaProperties()

	str, _ := String.ToString()
	arr, _ := Array.ToString()
	boolean, _ := Boolean.ToString()
	numeric, _ := Number.ToString()
	assert.Equal(t, len(properties), 5)
	assert.Equal(t, properties["username"].Type, str)
	assert.Equal(t, properties["email"].Type, str)
	assert.Equal(t, properties["array"].Type, arr)
	assert.Equal(t, properties["isMinor"].Type, boolean)
	assert.Equal(t, properties["age"].Type, numeric)

	schemas, err := userCreate.ToJsonSchema()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(schemas), 1)
}
