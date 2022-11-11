package grpc

import (
	"brinch/lib/utils"
	"fmt"
)

type Attribute struct {
	Type  string
	Name  string
	Index int
}

// ToProto generate a protobuf representation on the Attribute
func (attr *Attribute) ToProto() (string, error) {
	return fmt.Sprintf("%s %s = %d;", attr.Name, attr.Type, attr.Index), nil
}

// ToCode generate code representation for the Attribute
func (attr *Attribute) ToCode() (string, error) {
	return "", nil
}

func (attr *Attribute) FromStoredProcedureParameter(param *utils.StoredProcedureParameter, index int) (bool, error) {
	mapping, err := utils.ResolveTypeMappings(param.UdtName, param.Source, utils.Grpc)
	if err != nil {
		return false, err
	}

	attr.Type = mapping
	attr.Name = param.ParameterName
	attr.Index = index

	return true, nil
}
