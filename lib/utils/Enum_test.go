package utils

import (
	"brinch/lib/constants"
	"brinch/lib/utils/databases"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetListAllCustomTypesQuery(t *testing.T) {
	composite := Enums{}
	query := composite.GetQuery()

	if query != constants.ListEnums {
		t.Errorf("composite.GetQuery() = %s; want %s", query, constants.ListEnums)
	}
}

func TestGetEnumsSchema(t *testing.T) {
	t.Parallel()

	rows := pgxpoolmock.NewRows([]string{"enumtype", "enumlabel"}).
		AddRow("status", "draft").
		AddRow("status", "published").
		AddRow("channels", "web").
		AddRow("channels", "mobile").
		ToPgxRows()
	enum := Enums{}
	dbMeta := databases.DbMeta{}
	ok, _ := enum.QueryHandler(rows, &dbMeta)
	assert.Equal(t, ok, true)
	assert.Equal(t, len(enum.enums), 2)
	assert.Equal(t, enum.enums["status"][0], "draft")
	assert.Equal(t, enum.enums["channels"][0], "web")

	schemas, err := enum.ToJsonSchema()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(schemas), 2)
}
