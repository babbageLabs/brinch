package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babbageLabs/brinch/bin/core/methods"
	"github.com/babbageLabs/brinch/bin/core/transports"
	"github.com/babbageLabs/brinch/bin/static/queries/postgres"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func GetSp(t *testing.T) *StoredProcedures {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sp := &StoredProcedures{
		DB: db,
	}
	tr := &transports.MockTransport{}
	err = sp.SetTransport(tr)
	assert.NoError(t, err)

	columns := []string{"routine_schema", "specific_name", "routine_name", "parameter_name", "parameter_mode", "data_type", "udt_name", "parameter_default"}
	rows := sqlmock.NewRows(columns).
		AddRow("acl", "permissions_add_17880", "permissions_add", "permissions", "IN", "ARRAY", "_permission_tt", "").
		AddRow("acl", "permissions_add_17880", "permissions_add", "meta", "IN", "USER-DEFINED", "meta_tt", "").
		AddRow("acl", "role_permissions_add_17887", "role_permissions_add", "role_id", "INOUT", "bigint", "int8", "").
		AddRow("acl", "role_permissions_add_17887", "role_permissions_add", "meta", "IN", "USER-DEFINED", "meta_tt", "").
		AddRow("acl", "roles_add_17886", "roles_add", "roles", "INOUT", "ARRAY", "_role_tt", "")
	mock.ExpectQuery(regexp.QuoteMeta(postgres.ListAllSps)).WillReturnRows(rows)

	err = sp.Initialize(postgres.ListAllSps)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing mock sps", err)
	}

	return sp
}

// ########################################################
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

func TestRoute_GetReqParams(t *testing.T) {
	sps := GetSp(t)
	for s, route := range sps.GetRoutes() {
		if s == "permissions_add" {
			assert.Equal(t, 2, len(route.GetReqParams()))
		} else if s == "role_permissions_add" {
			assert.Equal(t, 2, len(route.GetReqParams()))
		} else if s == "roles_add" {
			assert.Equal(t, 1, len(route.GetReqParams()))
		}
	}

}

func TestRoute_GetResParams(t *testing.T) {
	sps := GetSp(t)
	for s, route := range sps.GetRoutes() {
		if s == "permissions_add" {
			assert.Equal(t, 0, len(route.GetResParams()))
		} else if s == "role_permissions_add" {
			assert.Equal(t, 1, len(route.GetResParams()))
		} else if s == "roles_add" {
			assert.Equal(t, 1, len(route.GetResParams()))
		}
	}
}

func TestSpAsCallable(t *testing.T) {
	sps := GetSp(t)
	permissionsAdd, err := sps.GetRoute("permissions_add")
	assert.NoError(t, err)

	_, err = methods.Call(permissionsAdd)
	assert.NoError(t, err)
}
