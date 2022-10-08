package utils

import (
	"database/sql"
	"github.com/spf13/viper"
	"path/filepath"
)

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

type StoredProcedure struct {
	Name       string
	Parameters []StoredProcedureParameter
}

type CustomTypeAttr struct {
	AttrName         string
	TypeName         string
	TypeCategory     string
	AttrTypeName     string
	AttrTypeCategory string
}

type CustomTypes struct {
	Name         string
	TypeCategory string
}

func (sp *StoredProcedure) ToJsonSchema() JSONSchemaBase {
	schema := viper.GetString("jsonSchema.schema")
	path := viper.GetString("jsonSchema.targetPath")

	id := filepath.Join(path, sp.Name+".schema"+".json")

	return JSONSchemaBase{
		Schema:      schema,
		Id:          id,
		Title:       sp.Name,
		Name:        sp.Name + ".schema" + ".json",
		Description: "An application route",
		Required:    sp.getRequiredProperties(),
		Properties:  sp.getProperties(),
		SchemaType:  "object",
	}
}

func (sp *StoredProcedure) getRequiredProperties() []string {
	var required []string
	if len(sp.Parameters) > 0 {
		for _, value := range sp.Parameters {
			if value.ParameterDefault.Valid == false {
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
