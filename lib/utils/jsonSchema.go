package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/ioutil"
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

func (jsonInput *JSONSchemaBase) ToFile() {
	dirPath := viper.GetString("jsonSchema.targetPath")
	port := viper.GetString("jsonSchema.server.port")
	jsonInput.AppendDefaultProperties()

	filePath := jsonInput.Id
	fmt.Printf("file path is %s\n", filePath)

	if dirPath == "" {
		path, err := os.Getwd()
		cobra.CheckErr(err)

		dirPath = filepath.Join(path, "api", "json")
	}

	if port != "" {
		jsonInput.Id = "http://localhost:" + port + "/" + jsonInput.Name
	}

	content, err := json.Marshal(jsonInput)
	cobra.CheckErr(err)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0660)
		cobra.CheckErr(err)
	}

	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		cobra.CheckErr(err)
	}
}

func (jsonInput *JSONSchemaBase) AppendDefaultProperties() {
	schema := viper.GetString("jsonSchema.schema")
	path := viper.GetString("jsonSchema.targetPath")

	jsonInput.Schema = schema
	jsonInput.Name = jsonInput.Id + ".schema" + ".json"
	jsonInput.Title = cases.Title(language.English, cases.Compact).String(jsonInput.Id)
	jsonInput.Id = filepath.Join(path, jsonInput.Name)

}

func (schemas Schemas) Export() {
	for _, schema := range schemas {
		schema.ToFile()
	}
}
