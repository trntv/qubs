package broker

import (
	"strconv"
	"time"
)

type delivery struct {
	attempt  uint64
	index    uint64
	queuedAt time.Time
}

type Message struct {
	payload      []byte
	broadcasting bool
	delivery     *delivery
}

func (m *Message) Id() string {
	return strconv.FormatUint(m.delivery.index, 10)
}

func (m *Message) Payload() []byte {
	return m.payload
}

func NewMessage(payload []byte, broadcasting bool) *Message {
	return &Message{
		payload:      payload,
		broadcasting: broadcasting,
		delivery: &delivery{
			queuedAt: time.Now(),
		},
	}
}
