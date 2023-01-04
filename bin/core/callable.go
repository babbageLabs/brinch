package core

import (
	"github.com/babbageLabs/brinch/bin/transports"
)

type IValidatable interface {
	IsValid() bool
	Validate() (bool, error)
}

type Params []IValidatable

type ICallable interface {
	Validate(params *Params) (bool, error)
	MustValidateRequest() bool
	MustValidateResponse() bool

	GetReqParams() Params
	GetResParams() Params

	GetSubject() (string, error) // ge the name of the topic in Nats
	GetTransport() (transports.ITransport, error)
}

func (params *Params) Marshal() ([]byte, error) {
	return nil, nil
}

func Call(callable ICallable) (*transports.Response, error) {
	transport, err := callable.GetTransport()
	if err != nil {
		return nil, err
	}

	subject, err := callable.GetSubject()
	if err != nil {
		return nil, err
	}

	params := callable.GetReqParams()
	if callable.MustValidateRequest() {
		_, err := callable.Validate(&params)
		if err != nil {
			return &transports.Response{}, err
		}
	}

	marshalled, err := params.Marshal()
	if err != nil {
		return nil, err
	}

	meta := &transports.MetaData{}
	res, err := transport.Exec(subject, marshalled, meta)
	if err != nil {
		return nil, err
	}

	if callable.MustValidateResponse() {
		params := callable.GetResParams()
		_, err := callable.Validate(&params)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func Validate(validatables []IValidatable) (bool, error) {
	for _, validatable := range validatables {
		isValid, err := validatable.Validate()
		if !isValid {
			return false, err
		}
	}

	return true, nil
}
