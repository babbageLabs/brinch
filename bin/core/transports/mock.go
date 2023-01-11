package transports

import (
	"github.com/babbageLabs/brinch/bin"
	"github.com/babbageLabs/brinch/bin/core/types"
)

type MockTransport struct {
}

func (mock *MockTransport) Close() (bool, error) {
	return true, nil
}

func (mock *MockTransport) Exec(subject string, msg []byte, meta *types.MetaData) (*types.Response, error) {
	bin.Logger.Debug("Mocking call to ", subject)
	return &types.Response{
		Data: msg,
		Meta: meta,
	}, nil

}

func (mock *MockTransport) Connect() (bool, error) {
	return true, nil
}
