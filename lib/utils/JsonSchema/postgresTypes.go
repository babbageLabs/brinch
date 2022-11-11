package JsonSchema

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SchemaTypeMapping struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type SchemaTypeMappingMap []SchemaTypeMapping

func (mapping *SchemaTypeMappingMap) ToMap() map[string]string {
	var DbTypeJsonSchemaMapping = make(map[string]string)
	for _, typeMapping := range *mapping {
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

	err = json.Unmarshal(body, &mapping)
	if err != nil {
		return false, err
	}

	return true, nil
}
