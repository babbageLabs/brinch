package utils

import (
	JsonSchema2 "brinch/lib/utils/JsonSchema"
	"brinch/lib/utils/databases"
	"database/sql"
)

type ParameterModes string

const (
	ParamIn    ParameterModes = "IN"
	ParamOut   ParameterModes = "OUT"
	ParamInOut ParameterModes = "INOUT"
)

// StoredProcedure stored procedures types
type StoredProcedure struct {
	Source     databases.SourceType
	Name       string
	Parameters []StoredProcedureParameter
}

// StoredProcedureParameter An Attribute belonging to a stored procedure as defined in the database
type StoredProcedureParameter struct {
	RoutineName      string
	RoutineSchema    string
	SpecificName     string
	ParameterName    string
	ParameterMode    ParameterModes
	DataType         string
	UdtName          string
	ParameterDefault sql.NullString
	Source           databases.SourceType
}

// CustomTypeAttr An Attribute belonging to a custom type as defined in the db
type CustomTypeAttr struct {
	AttrName         string
	TypeName         string
	TypeCategory     databases.PostgresTypeCategory
	AttrTypeName     string
	AttrTypeCategory databases.PostgresTypeCategory
	Source           databases.SourceType
}

// CustomTypeAttrs A slice of CustomTypeAttr with methods
type CustomTypeAttrs []CustomTypeAttr

func (sp *StoredProcedure) ToJsonSchema() JsonSchema2.JSONSchemaBase {
	return JsonSchema2.JSONSchemaBase{
		Id:          sp.Name,
		Description: "An application route",
		Required:    sp.getRequiredProperties(),
		Properties:  sp.getProperties(),
		SchemaType:  JsonSchema2.Object,
	}
}

func (sp *StoredProcedure) getRequiredProperties() []string {
	var required []string
	if len(sp.Parameters) > 0 {
		for _, value := range sp.Parameters {
			if !value.ParameterDefault.Valid {
				required = append(required, value.ParameterName)
			}
		}
	}
	return required
}

func (sp *StoredProcedure) getProperties() JsonSchema2.JSONSchemaProperties {
	var properties = JsonSchema2.JSONSchemaProperties{}

	for _, param := range sp.Parameters {
		dataType := param.DataType
		if dataType == "ARRAY" {
			dataType = param.UdtName[1:]
		} else {
			dataType = param.UdtName
		}

		properties[param.ParameterName] = JsonSchema2.JSONSchemaProperty{
			Type: dataType,
		}
	}

	return properties
}

// GetParameters return the input parameters for the SP
func (sp *StoredProcedure) GetParameters() []StoredProcedureParameter {
	var res []StoredProcedureParameter
	for _, parameter := range sp.Parameters {
		if parameter.ParameterMode != ParamOut {
			res = append(res, parameter)
		}
	}

	return res
}

// GetResponse return the expected response structure from this SP
func (sp *StoredProcedure) GetResponse() []StoredProcedureParameter {
	var res []StoredProcedureParameter
	for _, parameter := range sp.Parameters {
		if parameter.ParameterMode != ParamIn {
			res = append(res, parameter)
		}
	}

	return res
}

func (attrs CustomTypeAttrs) ToJsonSchemaProperties() JsonSchema2.JSONSchemaProperties {
	properties := make(JsonSchema2.JSONSchemaProperties)
	for _, attr := range attrs {
		properties[attr.AttrName] = JsonSchema2.JSONSchemaProperty{
			Type: string(attr.AttrTypeCategory),
		}
	}

	return properties

}
