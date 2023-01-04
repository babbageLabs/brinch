package transports

import "github.com/babbageLabs/brinch/bin"

type MockTransport struct {
}

func (mock *MockTransport) Close() (bool, error) {
	return true, nil
}

func (mock *MockTransport) Exec(subject string, msg []byte, meta *MetaData) (*Response, error) {
	bin.Logger.Debug("Mocking call to ", subject)
	return &Response{
		data: msg,
		meta: meta,
	}, nil

}

func (mock *MockTransport) Connect() (bool, error) {
	return true, nil
}
