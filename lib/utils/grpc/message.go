package grpc

import "fmt"

type message struct {
	name       string
	attributes []attribute
}

// ToProto generate a protobuf representation on the message
func (mes *message) ToProto() (string, error) {
	message := fmt.Sprintf("message %s { \n", mes.name)

	for _, a := range mes.attributes {
		attr, err := a.ToProto()
		if err != nil {
			return "", err
		}
		message = fmt.Sprintf("%s \n  %s", message, attr)
	}

	return fmt.Sprintf("%s \n  }", message), nil
}

// ToCode generate code representation for the message
func (mes *message) ToCode() (string, error) {
	return "", nil
}
