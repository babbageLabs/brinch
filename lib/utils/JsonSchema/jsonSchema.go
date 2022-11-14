package JsonSchema

import (
	"brinch/lib/utils/files"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
)

type SchemaType string
type Schemas []JSONSchemaBase

const (
	String  SchemaType = "string"
	Number  SchemaType = "number"
	Integer SchemaType = "integer"
	Object  SchemaType = "object"
	Array   SchemaType = "array"
	Boolean SchemaType = "boolean"
	Null    SchemaType = "null"
)

func (s SchemaType) ToString() (string, error) {
	switch s {
	case String:
		return "string", nil
	case Number:
		return "number", nil
	case Integer:
		return "integer", nil
	case Object:
		return "object", nil
	case Array:
		return "array", nil
	case Boolean:
		return "boolean", nil
	case Null:
		return "null", nil
	}
	return "", errors.New("unknown JSON schema type")
}

type JSONSchemaProperty struct {
	Type string `json:"type,omitempty"`
	Ref  string `json:"$ref,omitempty"`
}

type JSONSchemaProperties map[string]JSONSchemaProperty

type JSONSchemaBase struct {
	Schema      string               `json:"$schema" validate:"required"`
	Id          string               `json:"$id" validate:"required"`
	Title       string               `json:"title" validate:"required"`
	Name        string               `json:"-"`
	Description string               `json:"description"`
	SchemaType  SchemaType           `json:"type" validate:"required"`
	Required    []string             `json:"required,omitempty"`
	Enum        []string             `json:"enum,omitempty"`
	Properties  JSONSchemaProperties `json:"properties,omitempty"`
}

func (jsonInput *JSONSchemaBase) GetDirPath(dirPath string) (string, error) {
	if dirPath == "" {
		path, err := os.Getwd()
		if err != nil {
			return "", err
		}
		dirPath = filepath.Join(path, "api", "json")
	}

	return dirPath, nil
}

func (jsonInput *JSONSchemaBase) GetJsonServerHost(jsonServerHost string) string {
	if jsonServerHost != "" {
		jsonInput.Id = jsonServerHost + "/" + jsonInput.Name
	}

	return jsonInput.Id
}

func (jsonInput *JSONSchemaBase) ToBytes(jsonServerHost string) ([]byte, error) {
	jsonInput.AppendDefaultProperties()
	jsonInput.Id = jsonInput.GetJsonServerHost(jsonServerHost)

	return json.Marshal(jsonInput)
}

func (jsonInput *JSONSchemaBase) AppendDefaultProperties() {
	schema := viper.GetString("jsonSchema.schema")
	jsonInput.Schema = schema
	jsonInput.Name = jsonInput.Id + ".schema" + ".json"
	jsonInput.Title = cases.Title(language.English, cases.Compact).String(jsonInput.Id)
}

func (schemas *Schemas) Export(dirPath string, jsonServerHost string) {
	for _, schema := range *schemas {
		dirPath, err := schema.GetDirPath(dirPath)
		filePath := filepath.Join(dirPath, schema.Name)
		bytes, err := schema.ToBytes(jsonServerHost)
		if err != nil {
			logrus.Error(fmt.Sprintf("error preparing %s for writting to disk. the file will be skipped \n %s", schema.Id, err))
		} else {
			_, err := files.WriteToFile(filePath, bytes)
			if err != nil {
				logrus.Error(fmt.Sprintf("error writting %s to disk. \n %s", schema.Id, err))
			}
			logrus.Info(fmt.Sprintf("%s written to disk successfully", schema.Id))
		}
	}
}
