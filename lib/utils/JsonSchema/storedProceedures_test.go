package JsonSchema

import (
	"brinch/lib/constants"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetStoredProcedureQuery(t *testing.T) {
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
		AddRow("acl", "permissions_add_17880", "permissions_add", "permissions", "IN", "ARRAY", "_permission_tt", nil).
		ToPgxRows()

	sp := StoredProcedures{}
	schema := sp.QueryHandler(rows)

	assert.Equal(t, len(schema), 1)
}
