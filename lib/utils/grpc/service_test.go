package grpc

import (
	"brinch/lib/utils"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getRoute(index int) (Route, Route) {
	sps := getSps()
	route1 := Route{
		name: fmt.Sprintf("Route%d-1", index),
		sp:   sps.GetSps()[0],
	}

	route2 := Route{
		name: fmt.Sprintf("Route%d-2", index),
		sp:   sps.GetSps()[1],
	}

	return route1, route2
}

func TestService_ToCode(t *testing.T) {
	t.Parallel()
	var routes []Route
	route := Route{
		name: "",
		sp:   utils.StoredProcedure{},
	}

	service := Service{
		routes: append(routes, route),
	}

	code, err := service.ToCode()
	assert.Empty(t, err)
	assert.Empty(t, code)
}

func TestService_ToProto(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	// Exact URL match
	utils.GetMappingsUrl = "http://127.0.0.1"
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s/%s", utils.GetMappingsUrl, utils.Postgres, utils.Grpc),
		httpmock.NewStringResponder(200, `[{"key": "_permission_tt", "value": "value"}]`))

	var routes []Route
	r1, r2 := getRoute(1)
	r3, r4 := getRoute(2)
	r5, r6 := getRoute(3)
	routes = append(routes, r1, r2, r3, r4, r5, r6)

	options := make(map[string]string)

	base := BaseGRPC{
		Syntax:  "proto3",
		Options: options,
		Package: "package",
	}
	service := Service{
		routes,
		base,
	}

	proto, err := service.ToProto()
	assert.Empty(t, err)

	fmt.Printf("\n\n %s \n", proto)

}

func TestService_GetRouteMessagesProto(t *testing.T) {
	t.Parallel()
}
