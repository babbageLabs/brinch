package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var GetMappingsUrl = viper.GetString("app.typeMappings")

type SchemaTypeMapping struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type SchemaTypeMappingMap struct {
	mappings    []SchemaTypeMapping
	source      SourceType
	destination DestinationType
}

type SourceType string
type DestinationType string

const (
	Postgres SourceType = "postgres"
)

const (
	Grpc DestinationType = "grpc"
	//JsonSchema DestinationType = "jsonSchema"
)

func (mapping *SchemaTypeMappingMap) ToMap() map[string]string {
	var DbTypeJsonSchemaMapping = make(map[string]string)
	for _, typeMapping := range mapping.mappings {
		DbTypeJsonSchemaMapping[typeMapping.Key] = typeMapping.Value
	}

	return DbTypeJsonSchemaMapping
}

// GetMappings read key value mappings
func (mapping *SchemaTypeMappingMap) GetMappings(typeMappings string) (bool, error) {
	uri, err := url.ParseRequestURI(typeMappings)
	if err != nil {
		return false, err
	}
	resp, err := http.Get(uri.String())
	if err != nil {
		return false, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		cobra.CheckErr(err)
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(body, &mapping.mappings)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetSourceURl
func (mapping *SchemaTypeMappingMap) GetSourceURl() string {
	return fmt.Sprintf("%s/%s/%s", GetMappingsUrl, mapping.source, mapping.destination)
}

func ResolveTypeMappings(dbType string, source SourceType, destination DestinationType) (string, error) {
	// TODO optimize this function by introducing caching
	mapping := SchemaTypeMappingMap{source: source, destination: destination}
	_, err := mapping.GetMappings(mapping.GetSourceURl())
	if err != nil {
		return "", err
	}

	value, ok := mapping.ToMap()[dbType]
	if ok {
		return value, nil
	}

	return "", errors.New(fmt.Sprintf("Error resolving the type %s from %s to a  type in %s \n", dbType, source, destination))

}
