package transports

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babbageLabs/brinch/bin/core/database"
	"github.com/babbageLabs/brinch/bin/core/methods"
	"github.com/babbageLabs/brinch/bin/static/queries/postgres"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func GetSp(t *testing.T) *database.StoredProcedures {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sp := database.StoredProcedures{
		DB: db,
	}
	err = sp.SetTransport(&DBTransport{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when setting the transport", err)
	}

	columns := []string{"routine_schema", "specific_name", "routine_name", "parameter_name", "parameter_mode", "data_type", "udt_name", "parameter_default"}
	rows := sqlmock.NewRows(columns).
		AddRow("acl", "permissions_add_17880", "permissions_add", "permissions", "IN", "ARRAY", "_permission_tt", "").
		AddRow("acl", "permissions_add_17880", "permissions_add", "meta", "IN", "USER-DEFINED", "meta_tt", "")
	mock.ExpectQuery(regexp.QuoteMeta(postgres.ListAllSps)).WillReturnRows(rows)

	err = sp.Initialize(postgres.ListAllSps)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing mock sps", err)
	}

	return &sp
}

func TestDBTransportCallSpSuccess(t *testing.T) {
	sps := GetSp(t)
	addPermissions, err := sps.GetRoute("permissions_add")
	assert.NoError(t, err)
	_, err = methods.Call(addPermissions)
	assert.NoError(t, err) // TODO fix me
}
