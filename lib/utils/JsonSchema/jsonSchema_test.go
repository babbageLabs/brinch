package JsonSchema

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestJSONSchemaBase_AppendDefaultProperties(t *testing.T) {
	t.Parallel()
	schema := Base{
		Schema:      "https://json-schema.org/draft/2020-12/schema",
		Id:          "meta_tt",
		Name:        "",
		Description: "A composite type",
		SchemaType:  "object",
		Required:    nil,
		Enum:        nil,
		Properties:  nil,
	}

	schema.AppendDefaultProperties()

	assert.Equal(t, "Meta_tt", schema.Title)
	assert.Equal(t, "meta_tt.schema.json", schema.Name)

}

func TestJSONSchemaBase_GetDirPathDefault(t *testing.T) {
	t.Parallel()
	schema := Base{
		Schema:      "https://json-schema.org/draft/2020-12/schema",
		Id:          "meta_tt",
		Name:        "",
		Description: "A composite type",
		SchemaType:  "object",
		Required:    nil,
		Enum:        nil,
		Properties:  nil,
	}

	schema.AppendDefaultProperties()
	dirPath, err := schema.GetDirPath("")
	assert.NoError(t, err)

	path, err := os.Getwd()
	if err != nil {
		t.Errorf("could not read the current working directory")
	}

	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(path, "api", "json"), dirPath)

}

func TestJSONSchemaBase_GetJsonServerHost(t *testing.T) {
	t.Parallel()
	host := "http://localhost"
	schema := Base{
		Schema:      "https://json-schema.org/draft/2020-12/schema",
		Id:          "meta_tt",
		Name:        "",
		Description: "A composite type",
		SchemaType:  "object",
		Required:    nil,
		Enum:        nil,
		Properties:  nil,
	}
	schema.AppendDefaultProperties()
	id := schema.GetJsonServerHost(host)

	assert.Equal(t, host+"/"+schema.Name, id)
}

func TestJSONSchemaBase_ToBytes(t *testing.T) {
	t.Parallel()
	host := "http://localhost"
	schema := Base{
		Schema:      "https://json-schema.org/draft/2020-12/schema",
		Id:          "meta_tt",
		Name:        "",
		Description: "A composite type",
		SchemaType:  "object",
		Required:    nil,
		Enum:        nil,
		Properties:  nil,
	}

	_, err := schema.ToBytes(host)
	assert.NoError(t, err)
}
