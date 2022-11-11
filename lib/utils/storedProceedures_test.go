package utils

import (
	"brinch/lib/constants"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStoredProcedureQuery(t *testing.T) {
	t.Parallel()
	sp := StoredProcedures{}
	query := sp.GetQuery()

	if query != constants.ListAllSps {
		t.Errorf("sp.GetQuery() = %s; want %s", query, constants.ListAllCustomTypes)
	}
}

func TestStoredProcedureQueryHandler(t *testing.T) {
	t.Parallel()
	columns := []string{"specific_schema", "specific_name", "routine_name", "parameter_name", "parameter_mode", "data_type", "udt_name", "parameter_default"}
	rows := pgxpoolmock.
		NewRows(columns).
		AddRow("acl", "permissions_add_17880", "permissions_add", "permissions", ParamIn, "ARRAY", "_permission_tt", nil).
		AddRow("acl", "permissions_add_17380", "permissions_addD", "permissions", ParamIn, "ARRAY", "_permission_tt", nil).
		ToPgxRows()

	procedures := StoredProcedures{}
	dbMeta := DbMeta{}
	ok, err := procedures.QueryHandler(rows, &dbMeta)
	assert.Equal(t, err, nil)
	assert.Equal(t, ok, true)

	assert.Equal(t, len(procedures.sps), 2)

	schemas, err := procedures.ToJsonSchema()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(schemas), 2)
}
