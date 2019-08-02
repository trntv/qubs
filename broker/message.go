package broker

import "time"

type Message struct {
	Payload []byte

	tags         []string
	broadcasting bool
	index        uint64
	queuedAt     time.Time
}

func NewMessage(tags []string, payload []byte, broadcasting bool) *Message {
	return &Message{
		tags:         tags,
		Payload:      payload,
		broadcasting: broadcasting,
		queuedAt:     time.Now(),
	}
}
