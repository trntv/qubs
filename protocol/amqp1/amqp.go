// +build amqp

package amqp1

import (
	"errors"
	"github.com/trntv/qubs/broker"
	"qpid.apache.org/amqp"
)

func decodeAmqpMessage(msg amqp.Message) (*broker.Message, error) {
	body, ok := msg.Body().([]byte)
	if ok {
		return broker.NewMessage(make([]string, 0), body, false), nil
	}

	strBody, ok := msg.Body().(string)
	if ok {
		body = []byte(strBody)
		return broker.NewMessage(make([]string, 0), body, false), nil
	}

	return nil, errors.New("cannot conver AMQP message body")
}

func encodeAmqpMessage(msg *broker.Message) (amqp.Message, error) {
	am := amqp.NewMessageWith(msg.Payload())

	// @todo set properties

	return am, nil
}
