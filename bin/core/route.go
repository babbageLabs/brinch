package core

type Param struct {
	Name  string
	Value []byte
	Type  string
	err   error
}

type Route struct {
	Name             string
	Parameters       []Param
	Response         []Param
	validateRequest  bool
	validateResponse bool
}

// ########################## Params

// IsValid Check for validation error
func (param *Param) IsValid() bool {
	return param.err == nil
}

func (param *Param) Validate() (bool, error) {
	return false, nil
}

// ################################ Routes

func (route *Route) Validate(params []IValidatable) (bool, error) {
	return Validate(params)
}

func (route *Route) MustValidateRequest() bool {
	return route.validateRequest
}

func (route *Route) MustValidateResponse() bool {
	return route.validateResponse
}
