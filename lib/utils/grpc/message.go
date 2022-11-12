package grpc

import "fmt"

type Message struct {
	Name       string
	Attributes []Attribute
}

// ToProto generate a protobuf representation on the Message
func (mes *Message) ToProto() (string, error) {
	message := fmt.Sprintf("message %s { \n", mes.Name)

	for _, a := range mes.Attributes {
		attr, err := a.ToProto()
		if err != nil {
			return "", err
		}
		message = fmt.Sprintf("%s \n  %s", message, attr)
	}

	return fmt.Sprintf("%s\n}", message), nil
}

// ToCode generate code representation for the Message
func (mes *Message) ToCode() (string, error) {
	return "", nil
}
