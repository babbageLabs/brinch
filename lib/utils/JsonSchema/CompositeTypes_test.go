package JsonSchema

import (
	"brinch/lib/constants"
	"testing"
)

func TestGetQuery(t *testing.T) {
	composite := CompositeType{}
	query := composite.GetQuery()

	if query != constants.ListAllCustomTypes {
		t.Errorf("composite.GetQuery() = %s; want %s", query, constants.ListAllCustomTypes)
	}
}
