package grpc

import (
	"brinch/lib/utils"
	"database/sql"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttribute_ToCode(t *testing.T) {
	t.Parallel()

	attr := Attribute{
		Type:  "int",
		Name:  "username",
		Index: 1,
	}

	code, err := attr.ToCode()
	assert.Equal(t, code, "")
	assert.Equal(t, err, nil)
}

func TestAttribute_ToProto(t *testing.T) {
	t.Parallel()

	attr := Attribute{
		Type:  "int",
		Name:  "username",
		Index: 1,
	}

	code, err := attr.ToProto()
	assert.Equal(t, code, fmt.Sprintf("%s %s = %d;", attr.Name, attr.Type, attr.Index))
	assert.Equal(t, err, nil)
}

func TestAttribute_FromStoredProcedureParameter(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	attr := Attribute{}
	param := utils.StoredProcedureParameter{
		RoutineName:      "",
		RoutineSchema:    "",
		SpecificName:     "",
		ParameterName:    "",
		ParameterMode:    "",
		DataType:         "",
		UdtName:          "key",
		ParameterDefault: sql.NullString{},
		Source:           utils.Postgres,
	}

	// Exact URL match
	utils.GetMappingsUrl = "https://api.mybiz.com/articles"
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s/%s", utils.GetMappingsUrl, utils.Postgres, utils.Grpc),
		httpmock.NewStringResponder(200, `[{"key": "key", "value": "value"}]`))

	ok, err := attr.FromStoredProcedureParameter(&param, 1)
	assert.Empty(t, err)
	assert.Equal(t, ok, true)

	assert.Equal(t, attr.Name, param.ParameterName)
	assert.Equal(t, "value", attr.Type)
	assert.Equal(t, 1, attr.Index)
}
