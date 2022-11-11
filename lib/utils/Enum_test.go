package utils

import (
	"brinch/lib/constants"
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

	ok, _ := enum.QueryHandler(rows)
	assert.Equal(t, ok, true)
	assert.Equal(t, len(enum.enums), 2)
	assert.Equal(t, enum.enums["status"][0], "draft")
	assert.Equal(t, enum.enums["channels"][0], "web")
}
