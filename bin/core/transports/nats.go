package transports

import (
	"fmt"
	"github.com/babbageLabs/brinch/bin/core/types"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

type NatsMode int

const (
	Request NatsMode = iota
	Publish
)

type Nats struct {
	connection    *nats.Conn
	URL           string
	mode          NatsMode
	wg            sync.WaitGroup
	subscriptions map[string]*nats.Subscription
}

func (nts *Nats) Connect() (bool, error) {
	nc, err := nats.Connect(nts.URL,
		nats.DisconnectErrHandler(nts.DisconnectErrHandler),
		nats.ReconnectHandler(nts.ReconnectHandler),
		nats.ClosedHandler(nts.ClosedHandler),
		nats.ErrorHandler(nts.ErrorHandler),
	)
	if err != nil {
		return false, err
	}
	nts.connection = nc

	return true, nil
}

func (nts *Nats) Close() (bool, error) {
	// https://docs.nats.io/using-nats/developer/receiving/drain
	// Drain the connection, which will close it when done.
	if err := nts.connection.Drain(); err != nil {
		nts.connection.Close()
		return false, err
	}

	return true, nil
}

func (nts *Nats) Exec(subject string, msg []byte, meta *types.MetaData) (*types.Response, error) {
	switch nts.mode {
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

func (nts *Nats) DisconnectErrHandler(_ *nats.Conn, err error) {
	log.Printf("client disconnected: %v", err)
}

func (nts *Nats) ReconnectHandler(_ *nats.Conn) {
	log.Printf("client reconnected")
}

func (nts *Nats) ClosedHandler(_ *nats.Conn) {
	log.Printf("client closed")
}

func (nts *Nats) ErrorHandler(_ *nats.Conn, s *nats.Subscription, err error) {
	if s != nil {
		log.Printf("Async error in %q/%q: %v", s.Subject, s.Queue, err)
	} else {
		log.Printf("Async error outside subscription: %v", err)
	}
}

func (nts *Nats) Publish(subject string, msg []byte) (bool, error) {
	err := nts.connection.Publish(subject, msg)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (nts *Nats) Subscribe(subj string, cb nats.MsgHandler) error {
	sub, err := nts.connection.Subscribe(subj, cb)
	if err != nil {
		return err
	}

	if nts.subscriptions != nil {
		nts.subscriptions[subj] = sub
	} else {
		nts.subscriptions = map[string]*nats.Subscription{subj: sub}
	}

	return nil
}

func (nts *Nats) QueueSubscribe(subj string, queue string, cb nats.MsgHandler) error {
	sub, err := nts.connection.QueueSubscribe(subj, queue, cb)
	if err != nil {
		return err
	}

	if nts.subscriptions != nil {
		nts.subscriptions[subj] = sub
	} else {
		nts.subscriptions = map[string]*nats.Subscription{subj: sub}
	}

	return nil
}

func (nts *Nats) UnSubscribe(subj string) (bool, error) {
	subscription, ok := nts.subscriptions[subj]

	if ok {
		// Call Drain on the subscription. It unsubscribes but
		// wait for all pending messages to be processed.
		if err := subscription.Drain(); err != nil {
			return false, err
		}
	} else {
		return false, fmt.Errorf(fmt.Sprintf("subscription %s was not found", subj))
	}

	return true, nil
}
