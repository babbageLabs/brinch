package grpc

import (
	"brinch/lib/utils"
	"brinch/lib/utils/databases"
	"fmt"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getSps() utils.StoredProcedures {
	columns := []string{"specific_schema", "specific_name", "routine_name", "parameter_name", "parameter_mode", "data_type", "udt_name", "parameter_default"}
	rows := pgxpoolmock.
		NewRows(columns).
		AddRow("acl", "permissions_add_17880", "permissions_add", "permissions", utils.ParamIn, "ARRAY", "_permission_tt", nil).
		AddRow("acl", "permissions_add_17380", "permissions_addD", "permissions", utils.ParamOut, "ARRAY", "_permission_tt", nil).
		ToPgxRows()

	procedures := utils.StoredProcedures{}
	dbMeta := databases.DbMeta{
		SourceType: databases.Postgres,
	}
	_, _ = procedures.QueryHandler(rows, &dbMeta)

	return procedures
}

func TestRoute_ToCode(t *testing.T) {
	t.Parallel()

	mes := Route{
		name: "login",
	}

	code, err := mes.ToCode()
	assert.Empty(t, code, "")
	assert.Equal(t, err, nil)
}

func TestRoute_ToProto(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	sps := getSps()
	route := Route{
		name: "permissions_add",
		sp:   sps.GetSps()[0],
	}
	// Exact URL match
	utils.GetMappingsUrl = "https://api.mybiz.com/articles"
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s/%s", utils.GetMappingsUrl, databases.Postgres, utils.Grpc),
		httpmock.NewStringResponder(200, `[{"key": "_permission_tt", "value": "value"}]`))

	proto, err := route.ToProto()
	assert.Empty(t, err)

	fmt.Printf("\n\n %s \n", proto)
}

func TestRoute_GetMessages(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	sps := getSps()
	route := Route{
		name: "permissions_add",
		sp:   sps.GetSps()[0],
	}
	// Exact URL match
	utils.GetMappingsUrl = "https://api.mybiz.com/articles"
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s/%s", utils.GetMappingsUrl, databases.Postgres, utils.Grpc),
		httpmock.NewStringResponder(200, `[{"key": "_permission_tt", "value": "value"}]`))

	messages := route.getMessages()
	assert.Equal(t, 2, len(messages))

	fmt.Printf("\n\n %v \n", messages)
}

func TestRoute_GetMessagesError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	sps := getSps()
	route := Route{
		name: "permissions_add",
		sp:   sps.GetSps()[0],
	}
	// Exact URL match
	utils.GetMappingsUrl = "https://api.mybiz.com/articles"
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s/%s", utils.GetMappingsUrl, utils.JsonSchema, utils.Grpc),
		httpmock.NewStringResponder(200, `[{"key": "_permission_tt", "value": "value"}]`))

	messages := route.getMessages()
	assert.Equal(t, []Message(nil), messages)

	fmt.Printf("\n\n %v \n", messages)
}

func TestRoute_GetMessagesProto(t *testing.T) {
	t.Parallel()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	sps := getSps()
	route := Route{
		name: "permissions_add",
		sp:   sps.GetSps()[0],
	}
	// Exact URL match
	utils.GetMappingsUrl = "https://api.mybiz.com/articles"
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s/%s", utils.GetMappingsUrl, databases.Postgres, utils.Grpc),
		httpmock.NewStringResponder(200, `[{"key": "_permission_tt", "value": "value"}]`))

	messages, err := route.GetMessagesProto()
	assert.Empty(t, err)

	fmt.Printf("\n\n %s \n", messages)
}
