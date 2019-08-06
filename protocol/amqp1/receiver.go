// +build amqp

package amqp1

import (
	"github.com/sirupsen/logrus"
	"github.com/trntv/qubs/broker"
	"qpid.apache.org/proton"
)

// amqpReceiver has a channel to buffer messages that have been received by the
// AmqpConnectionHandler and are waiting to go on the queue. AMQP credit ensures that the
// AmqpConnectionHandler does not overflow the buffer and block.
type amqpReceiver struct {
	l      proton.Link
	h      *AmqpConnectionHandler
	buffer chan receivedMessage
	queue  *broker.Queue
}

// run runs in a separate goroutine. It moves messages from the buffer to the
// queue for a amqpReceiver link, and injects a AmqpConnectionHandler function to acknowledge the
// message and send a credit.
func (r *amqpReceiver) run() {
	for rm := range r.buffer {
		d := rm.delivery
		msg, err := decodeAmqpMessage(rm.message)
		if err != nil {
			logrus.Error(err)
			continue
		}

		// We are not in the AmqpConnectionHandler goroutine so we Inject the Accept function as a closure.
		err = r.h.injecter.Inject(func() {
			r.queue.Deliver(msg)
			if r == r.h.receivers[r.l] {
				d.Accept()  // Accept the delivery
				r.l.Flow(1) // Add one credit
			}
		})
		if err != nil {
			logrus.Error(err)
		}
	}
}

// stop closes the buffer channel and waits for the run() goroutine to stop.
func (r *amqpReceiver) stop() {
	close(r.buffer)
}
