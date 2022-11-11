package grpc

import "fmt"

type BaseGRPC struct {
	Syntax  string
	Options map[string]string
	Name    string
	Package string
}

// GetDefaultProtoAttributes generate a protobuf representation on the grpc item
func (grpc *BaseGRPC) GetDefaultProtoAttributes() (string, error) {
	proto := ""
	// add the syntax attr
	if grpc.Syntax != "" {
		proto = fmt.Sprintf("syntax = \"%s\";", proto)
	}

	// add service options
	for k, v := range grpc.Options {
		proto = fmt.Sprintf("%s \n option %s = \"%s\";", proto, k, v)
	}

	//add package
	if grpc.Package != "" {
		proto = fmt.Sprintf("%s \n\n package %s", proto, grpc.Package)
	}

	proto = fmt.Sprintf("%s \n", proto)

	return proto, nil
}
