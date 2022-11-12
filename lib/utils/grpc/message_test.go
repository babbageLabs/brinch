package grpc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_ToCode(t *testing.T) {
	t.Parallel()

	mes := Message{
		Name: "username",
	}

	code, err := mes.ToCode()
	assert.Equal(t, code, "")
	assert.Equal(t, err, nil)
}

func TestMessage_ToProto(t *testing.T) {
	t.Parallel()
	var attrs []Attribute

	attrs = append(attrs, Attribute{
		Type:  "string",
		Name:  "username",
		Index: 1,
	})

	attrs = append(attrs, Attribute{
		Type:  "string",
		Name:  "password",
		Index: 2,
	})

	mes := Message{
		Name:       "Login",
		Attributes: attrs,
	}

	code, err := mes.ToProto()
	assert.Equal(t, err, nil)

	fmt.Printf("\n\n %s \n", code)
}
