package grpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseGRPC_GetDefaultProtoAttributes(t *testing.T) {
	t.Parallel()
	options := make(map[string]string)
	options["java_outer_classname"] = "RouteGuideProto"
	options["go_package"] = "google.golang.org/grpc/examples/route_guide/routeguide"

	rpc := BaseGRPC{
		Syntax:  "proto3",
		Options: options,
		Name:    "",
		Package: "routeguide",
	}

	_, err := rpc.GetDefaultProtoAttributes()
	assert.Empty(t, err)
}
