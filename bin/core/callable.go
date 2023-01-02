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

	GetReqParams() *Params
	GetResParams() *Params

	GetSubject() string // ge the name of the topic in Nats
	transports.Transport
}

func (params *Params) Marshal() ([]byte, error) {
	return nil, nil
}

func Call(callable ICallable) (*transports.Response, error) {
	if callable.MustValidateRequest() {
		_, err := callable.Validate(callable.GetReqParams())
		if err != nil {
			return &transports.Response{}, err
		}
	}

	marshalled, err := callable.GetReqParams().Marshal()
	if err != nil {
		return &transports.Response{}, err
	}

	meta := &transports.MetaData{}
	res, err := callable.Exec(callable.GetSubject(), marshalled, meta)
	if err != nil {
		return &transports.Response{}, err
	}

	if callable.MustValidateResponse() {
		_, err := callable.Validate(callable.GetResParams())
		if err != nil {
			return &transports.Response{}, err
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
