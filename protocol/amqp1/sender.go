// +build amqp

package amqp1

import (
	"github.com/trntv/qubs/broker"
	"qpid.apache.org/proton"
)

type sendResult struct {
	err error
}

// amqpSender has a channel that is used to signal when there is credit to send messages.
type amqpSender struct {
	h *AmqpConnectionHandler
	l proton.Link
	q *broker.Queue
	s *broker.Consumer

	mch    chan broker.Message
	rch    chan sendResult
	closed bool
}

func (s *amqpSender) start() {
	for msg := range s.mch {
		s.h.injecter.Inject(func() { // Inject handler function to actually send
			am, _ := encodeAmqpMessage(&msg)
			delivery, err := s.l.Send(am)
			if err == nil {
				delivery.Settle()
			}

			s.rch <- sendResult{err: err}
		})
	}

}

func (s *amqpSender) Send(msg broker.Message) error {
	s.mch <- msg

	select {
	case r := <-s.rch:
		return r.err
	}
}

// stop closes the credit channel and waits for the run() goroutine to stop.
func (s *amqpSender) Close() error {
	close(s.mch)
	s.closed = true
	return nil
}

func (s *amqpSender) IsClosed() bool {
	return s.closed
}
