package transports

import (
	"fmt"
	"github.com/babbageLabs/brinch/bin/core/types"
	"github.com/nats-io/nats.go"
	"time"
)

type NatsMode int

const (
	Request NatsMode = iota
	Publish
)

type Nats struct {
	connection *nats.Conn
	URL        string
	mode       NatsMode
}

func (nts *Nats) Connect() (bool, error) {
	nc, err := nats.Connect(nts.URL)
	if err != nil {
		return false, err
	}
	nts.connection = nc

	return true, nil
}

func (nts *Nats) Close() (bool, error) {
	// https://docs.nats.io/using-nats/developer/receiving/drain
	if err := nts.connection.Drain(); err != nil {
		nts.connection.Close()
	}

	return true, nil
}

func (nts *Nats) Request(subject string, msg []byte) (*nats.Msg, error) {
	res, err := nts.connection.Request(subject, msg, 30*time.Second)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (nts *Nats) Publish(subject string, msg []byte) (bool, error) {
	err := nts.connection.Publish(subject, msg)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (nts *Nats) Exec(subject string, msg []byte, meta *types.MetaData) (*types.Response, error) {
	switch nts.mode {
	case Request:
		res, err := nts.Request(subject, msg)
		if err != nil {
			return &types.Response{}, err
		}
		return &types.Response{
			Meta: meta,
			Data: res.Data,
		}, nil
	case Publish:
		_, err := nts.Publish(subject, msg)
		if err != nil {
			return nil, err
		}
		return &types.Response{}, nil
	default:
		return nil, fmt.Errorf("transport.NATS.mode is not configured")
	}
}
