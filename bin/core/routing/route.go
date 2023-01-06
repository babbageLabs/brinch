package routing

import (
	"fmt"
	"github.com/babbageLabs/brinch/bin/core/methods"
	"github.com/babbageLabs/brinch/bin/core/types"
)

const (
	IN    types.ParameterMode = "IN"
	OUT   types.ParameterMode = "OUT"
	INOUT types.ParameterMode = "INOUT"
)

type Param struct {
	Name    string
	Value   []byte
	Type    string
	Err     error
	Mode    types.ParameterMode
	IsArray bool
}

type Route struct {
	NameSpace        string
	Name             string
	Parameters       []types.Param
	ValidateRequest  bool
	ValidateResponse bool
	Transport        *types.ITransport
}

func (route *Route) GetName() (string, error) {
	return route.Name, nil
}

func (route *Route) AddParam(param types.Param) (bool, error) {
	for _, parameter := range route.Parameters {
		name, err := param.GetName()
		if err != nil {
			return false, err
		}
		mode, err := param.GetMode()
		if err != nil {
			return false, err
		}

		pname, err := parameter.GetName()
		if err != nil {
			return false, err
		}
		pmode, err := parameter.GetMode()
		if err != nil {
			return false, err
		}

		if pname == name && mode == pmode {
			return false, err
		}
	}
	route.Parameters = append(route.Parameters, param)

	return true, nil
}

func (route *Route) Validate(params *types.IValidatable) (bool, error) {
	var p types.Validatables
	p = append(p, *params)
	return methods.Validate(p)
}

func (route *Route) GetReqParams() types.Params {
	var params types.Params
	for _, parameter := range route.Parameters {
		if parameter.IsInput() {
			params = append(params, parameter)
		}
	}

	return params
}

func (route *Route) GetResParams() types.Params {
	var params types.Params
	for _, parameter := range route.Parameters {
		if parameter.IsOutPut() {
			params = append(params, parameter)
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

func (route *Route) GetTransport() (types.ITransport, error) {
	if route.Transport == nil {
		return nil, fmt.Errorf("message transport is not specified")
	}
	return *route.Transport, nil
}

func (route *Route) MustValidateRequest() bool {
	return route.ValidateRequest
}

func (route *Route) MustValidateResponse() bool {
	return route.ValidateResponse
}

func (route *Route) SetNameSpace(name string) error {
	route.NameSpace = name

	return nil
}

// ########################## Params

// IsValid Check for validation error
func (param *Param) IsValid() bool {
	return param.Err == nil
}

func (param *Param) Validate() (bool, error) {
	// TODO implement this
	//panic("implement this")
	return false, nil
}

func (param *Param) IsInput() bool {
	return param.Mode == IN || param.Mode == INOUT
}

func (param *Param) IsOutPut() bool {
	return param.Mode == OUT || param.Mode == INOUT
}

func (param *Param) GetName() (string, error) {
	if param.Name == "" {
		return "", fmt.Errorf("route name not set")
	}

	return param.Name, nil
}

func (param *Param) GetMode() (types.ParameterMode, error) {
	if param.Mode == "" {
		return "", fmt.Errorf("route mode not set")
	}

	return param.Mode, nil
}

func (param *Param) Marshal() ([]byte, error) {
	return nil, nil
}

// ################################ StoredProcedures

//func (route *Route) Validate(params []IValidatable) (bool, error) {
//	return Validate(params)
//}
