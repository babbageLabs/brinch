package core

import (
	"fmt"
	"github.com/babbageLabs/brinch/bin/transports"
)

type ParameterMode string

const (
	IN  ParameterMode = "IN"
	OUT ParameterMode = "OUT"
	//INOUT ParameterMode = "INOUT"
)

type Param struct {
	Name    string
	Value   []byte
	Type    string
	err     error
	Mode    ParameterMode
	isArray bool
}

type Route struct {
	NameSpace        string
	Name             string
	Parameters       []Param
	validateRequest  bool
	validateResponse bool
	transport        transports.ITransport
}

func (route *Route) Validate(params *Params) (bool, error) {
	return Validate(*params)
}

func (route *Route) GetReqParams() Params {
	var params Params
	for _, parameter := range route.Parameters {
		if parameter.IsInput() {
			params = append(params, &parameter)
		}
	}

	return params
}

func (route *Route) GetResParams() Params {
	var params Params
	for _, parameter := range route.Parameters {
		if parameter.IsOutPut() {
			params = append(params, &parameter)
		}
	}

	return params
}

func (route *Route) GetSubject() (string, error) {
	if route.NameSpace != "" {
		return fmt.Sprintf(route.NameSpace, ".", route.Name), nil
	}
	return "", fmt.Errorf("specify the namespace")
}

func (route *Route) GetTransport() (transports.ITransport, error) {
	if route.transport == nil {
		return nil, fmt.Errorf("message transport is not specified")
	}
	return route.transport, nil
}

// ########################## Params

// IsValid Check for validation error
func (param *Param) IsValid() bool {
	return param.err == nil
}

func (param *Param) Validate() (bool, error) {
	// TODO implement this
	//panic("implement this")
	return false, nil
}

func (param *Param) IsInput() bool {
	return param.Mode != OUT
}

func (param *Param) IsOutPut() bool {
	return param.Mode != IN
}

// ################################ StoredProcedures

//func (route *Route) Validate(params []IValidatable) (bool, error) {
//	return Validate(params)
//}

func (route *Route) MustValidateRequest() bool {
	return route.validateRequest
}

func (route *Route) MustValidateResponse() bool {
	return route.validateResponse
}

func (route *Route) AddParam(param *Param) bool {
	for _, parameter := range route.Parameters {
		if parameter.Name == param.Name && param.Mode == parameter.Mode {
			return false
		}
	}
	route.Parameters = append(route.Parameters, *param)

	return true
}
