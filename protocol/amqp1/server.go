// +build amqp

package amqp1

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trntv/qubs/broker"
	"net"
	"qpid.apache.org/amqp"
	"qpid.apache.org/proton"
	"time"
)

// receivedMessage holds a message and a Delivery so that the message can be
// acknowledged when it is put on the queue.
type receivedMessage struct {
	delivery proton.Delivery
	message  amqp.Message
}

type AmqpConnectionHandler struct {
	broker    *broker.Broker
	receivers map[proton.Link]*amqpReceiver
	senders   map[proton.Link]*amqpSender
	injecter  proton.Injecter
}

func NewAmqpServer(b *broker.Broker) func(net.Listener) {
	handler := NewHandler(b)
	return func(listener net.Listener) {
		for {
			conn, err := listener.Accept()
			if err != nil {
				logrus.Error(err)
				continue
			}

			err = conn.SetReadDeadline(time.Now().Add(30 * time.Second))
			if err != nil {
				logrus.Error(err)
				continue
			}

			go handler(conn)
		}
	}
}

func NewHandler(b *broker.Broker) func(conn net.Conn) {
	return func(conn net.Conn) {
		handler := &AmqpConnectionHandler{
			broker:    b,
			receivers: make(map[proton.Link]*amqpReceiver),
			senders:   make(map[proton.Link]*amqpSender),
		}

		adapter := proton.NewMessagingAdapter(handler)
		adapter.Prefetch = 0
		adapter.AutoAccept = false

		engine, err := proton.NewEngine(conn, adapter)
		if err != nil {
			logrus.Error(err)
			return
		}
		engine.Server() // Enable server-side protocol negotiation.
		logrus.Debugf("Accepted connection %s", engine)
		go func() { // Start goroutine to run the engine event loop
			err := engine.Run()
			if err != nil {
				logrus.Error(err)
			}
			logrus.Debugf("Closed")
		}()
	}
}

// HandleMessagingEvent handles an event, called in the AmqpConnectionHandler goroutine.
func (h *AmqpConnectionHandler) HandleMessagingEvent(t proton.MessagingEvent, e proton.Event) {
	switch t {

	case proton.MStart:
		h.injecter = e.Injecter()

	case proton.MLinkOpening:
		fmt.Println("Opening", e.Link())
		if e.Link().IsReceiver() {
			q := h.broker.GetQueue(e.Connection().Hostname(), e.Link().RemoteTarget().Address())
			r := &amqpReceiver{
				l:      e.Link(),
				h:      h,
				buffer: make(chan receivedMessage, 128),
				queue:  q,
			}
			h.receivers[r.l] = r
			r.l.Flow(128) // Give credit to fill the buffer to capacity.
			go r.run()
		} else {
			q := h.broker.GetQueue(e.Connection().Hostname(), e.Link().RemoteSource().Address())
			as := &amqpSender{
				l:   e.Link(),
				q:   q,
				h:   h,
				mch: make(chan broker.Message, 128),
				rch: make(chan sendResult, 1),
			}
			h.senders[e.Link()] = as
			go as.start()
		}

	case proton.MLinkClosed:
		h.linkClosed(e.Link(), e.Link().RemoteCondition().Error())

	case proton.MSendable:
		fmt.Println("Sendable!", e.Link())
		if as, ok := h.senders[e.Link()]; ok {
			as.s = as.q.AddConsumer(as, make([]string, 0))
		} else {
			proton.CloseError(e.Link(), amqp.Errorf(amqp.NotFound, "link %s amqpSender not found", e.Link()))
		}

	case proton.MMessage:
		m, err := e.Delivery().Message() // Message() must be called while handling the MMessage event.
		if err != nil {
			proton.CloseError(e.Link(), err)
			break
		}
		r, ok := h.receivers[e.Link()]
		if !ok {
			proton.CloseError(e.Link(), amqp.Errorf(amqp.NotFound, "link %s amqpReceiver not found", e.Link()))
			break
		}
		// This will not block as AMQP credit is set to the buffer capacity.
		r.buffer <- receivedMessage{e.Delivery(), m}
		logrus.Debugf("link %s received %#v", e.Link(), m)

	case proton.MConnectionClosed:
		fmt.Println("Closed", e.Link())
		for l, _ := range h.receivers {
			h.linkClosed(l, nil)
		}
		for l, _ := range h.senders {
			h.linkClosed(l, nil)
		}

	case proton.MDisconnected:
		fmt.Println("Disconnected", e.Transport().Condition().Error())
	}
}

// linkClosed is called when a link has been closed by both ends.
// It removes the link from the handlers maps and stops its goroutine.
func (h *AmqpConnectionHandler) linkClosed(l proton.Link, err error) {
	fmt.Println("closing", l)
	if s, ok := h.senders[l]; ok {
		s.Close()
		delete(h.senders, l)
	} else if r, ok := h.receivers[l]; ok {
		r.stop()
		delete(h.receivers, l)
	}
}
