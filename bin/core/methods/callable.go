package methods

import (
	"github.com/babbageLabs/brinch/bin/core/types"
)

func Call(callable types.ICallable) (*types.Response, error) {
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
		_, err := Validate(params)
		if err != nil {
			return nil, err
		}
	}

	msg, err := Marshal(&params)
	if err != nil {
		return nil, err
	}

	meta := &types.MetaData{}
	res, err := transport.Exec(subject, msg, meta)
	if err != nil {
		return res, err
	}

	if callable.MustValidateResponse() {
		params := callable.GetResParams()
		_, err := Validate(params)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
