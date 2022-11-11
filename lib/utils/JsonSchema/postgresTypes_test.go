package JsonSchema

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGetMapAndToMapInvalidURl(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mapping := SchemaTypeMappingMap{}

	uri := "/articles"
	_, err := mapping.GetMappings(uri)
	if err == nil {
		t.Errorf("Expected invalid url error. found no error")
	}
}

func TestGetMapAndToMapTestValidUrl(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mapping := SchemaTypeMappingMap{}

	uri := "https://api.mybiz.com/articles"
	// Exact URL match
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `[{"key": "key", "value": "value"},{"key": "key2", "value": "value2"}]`))

	_, err := mapping.GetMappings(uri)
	if err != nil {
		t.Errorf("Expected GetMappings to complete with no errors: %v", err)
	}

	if len(mapping) != 2 {
		t.Errorf("Expected type mapping length 2, got %d", len(mapping))
	}

	mappingMap := mapping.ToMap()
	if len(mappingMap) != 2 {
		t.Errorf("Expected type mappingMap length 2, got %d", len(mappingMap))
	}

	if mappingMap["key"] != "value" {
		t.Errorf("Expected value got %s", mappingMap["key"])
	}
}
