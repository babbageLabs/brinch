package grpc

import (
	"brinch/lib/utils"
	"fmt"
)

type Route struct {
	name     string
	sp       utils.StoredProcedure
	messages []message
}

// ToProto generate a protobuf representation on the route
func (route *Route) ToProto() (string, error) {
	parameters, err := route.GetParamType()
	if err != nil {
		return "", err
	}
	returnType, err := route.GetReturnType()
	if err != nil {
		return "", err
	}
	name := fmt.Sprintf("rpc %s(%s) returns (%s) {}", route.name, parameters, returnType)
	return name, nil
}

// ToCode generate code representation for the route
func (route *Route) ToCode() (string, error) {
	return "", nil
}

// GetMessagesProto get the input parameters to the route and return response for the route
func (route *Route) GetMessagesProto() (string, error) {
	messages := ""

	for _, m := range route.messages {
		msg, err := m.ToProto()
		if err != nil {
			return "", err
		}
		messages = fmt.Sprintf("%s \n  %s", messages, msg)
	}
	return messages, nil
}

// GetParamType Get the endpoint params
func (route *Route) GetParamType() (string, error) {
	return "", nil
}

// GetReturnType Get the endpoint return type
func (route *Route) GetReturnType() (string, error) {
	return "", nil
}
