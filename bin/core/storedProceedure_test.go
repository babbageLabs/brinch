package core

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babbageLabs/brinch/bin/static/queries/postgres"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestStoredProcedures_Initialize(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sp := &StoredProcedures{
		routes:    nil,
		sps:       nil,
		DB:        db,
		transport: nil,
	}

	columns := []string{"routine_schema", "specific_name", "routine_name", "parameter_name", "parameter_mode", "data_type", "udt_name", "parameter_default"}
	rows := sqlmock.NewRows(columns).
		AddRow("acl", "permissions_add_17880", "permissions_add", "permissions", "IN", "ARRAY", "_permission_tt", "").
		AddRow("acl", "permissions_add_17880", "permissions_add", "meta", "IN", "USER-DEFINED", "meta_tt", "").
		AddRow("acl", "role_permissions_add_17887", "role_permissions_add", "role_id", "IN", "bigint", "int8", "").
		AddRow("acl", "role_permissions_add_17887", "role_permissions_add", "meta", "IN", "USER-DEFINED", "meta_tt", "").
		AddRow("acl", "roles_add_17886", "roles_add", "roles", "IN", "ARRAY", "_role_tt", "")

	mock.ExpectQuery(regexp.QuoteMeta(postgres.ListAllSps)).WillReturnRows(rows)

	err = sp.Initialize(postgres.ListAllSps)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(sp.sps))
}
