package utils

import (
	"database/sql"
)

// StoredProcedure stored procedures types
type StoredProcedure struct {
	Name       string
	Parameters []StoredProcedureParameter
}

func (sp *StoredProcedure) ToJsonSchema() JSONSchemaBase {
	return JSONSchemaBase{
		Id:          sp.Name,
		Description: "An application route",
		Required:    sp.getRequiredProperties(),
		Properties:  sp.getProperties(),
		SchemaType:  Object,
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

func (sp *StoredProcedure) getProperties() JSONSchemaProperties {
	var properties = JSONSchemaProperties{}

	for _, param := range sp.Parameters {
		dataType := param.DataType
		if dataType == "ARRAY" {
			dataType = param.UdtName[1:]
		} else {
			dataType = param.UdtName
		}

		properties[param.ParameterName] = JSONSchemaProperty{
			Type: dataType,
		}
	}

	return properties
}

// GetParameters return the input parameters for the SP
func (sp *StoredProcedure) GetParameters() {

}

// GetResponse return the expected response structure from this SP
func (sp *StoredProcedure) GetResponse() {

}

// StoredProcedureParameter An Attribute belonging to a stored procedure as defined in the database
type StoredProcedureParameter struct {
	RoutineName      string
	RoutineSchema    string
	SpecificName     string
	ParameterName    string
	ParameterMode    string
	DataType         string
	UdtName          string
	ParameterDefault sql.NullString
}

// CustomTypeAttr An Attribute belonging to a custom type as defined in the db
type CustomTypeAttr struct {
	AttrName         string
	TypeName         string
	TypeCategory     PostgresTypeCategory
	AttrTypeName     string
	AttrTypeCategory PostgresTypeCategory
}

// CustomTypeAttrs A slice of CustomTypeAttr with methods
type CustomTypeAttrs []CustomTypeAttr

func (attrs CustomTypeAttrs) ToJsonSchemaProperties() JSONSchemaProperties {
	properties := make(JSONSchemaProperties)
	for _, attr := range attrs {
		properties[attr.AttrName] = JSONSchemaProperty{
			Type: string(attr.AttrTypeCategory.ToJsonType()),
		}
	}

	return properties

}
