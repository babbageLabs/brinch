package grpc

import "fmt"

type attribute struct {
	Type  string
	name  string
	index int
}

// ToProto generate a protobuf representation on the attribute
func (attr *attribute) ToProto() (string, error) {
	return fmt.Sprintf("%s %s = %d;", attr.name, attr.Type, attr.index), nil
}

// ToCode generate code representation for the attribute
func (attr *attribute) ToCode() (string, error) {
	return "", nil
}
