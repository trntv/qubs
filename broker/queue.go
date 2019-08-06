package broker

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Queue interface {
	Deliver(msg *Message) error
	Consume(fn func(*Message) error)
}

type queue struct {
	seq uint64

	messages  chan *Message
	consumers []*Consumer

	m_lock sync.Mutex
	c_lock sync.Mutex
}

func NewQueue() *queue {
	q := &queue{
		messages: make(chan *Message, 65536),
	}

	return q
}

func (q *queue) Deliver(msg *Message) error {
	for {
		cc := len(q.consumers)
		if cc == 0 {
			return q.Enqueue(msg)
		}

		seq := atomic.AddUint64(&q.seq, 1)
		i := seq % uint64(len(q.consumers))
		if q.consumers[i].link.IsClosed() {
			q.removeConsumer(i)
			continue
		}

		err := q.consumers[i].link.Send(msg)
		if err != nil {
			q.consumers[i].link.IsClosed()
			return q.Enqueue(msg)
		}

		return nil
	}
}

func (q *queue) Consume(fn func(*Message) error) {
	msg := q.Dequeue()
	err := fn(msg)
	if err != nil {
		_ = q.Enqueue(msg)
	}
}

func (q *queue) Enqueue(msg *Message) error {
	select {
	case q.messages <- msg:
	default:
		return errors.New("queue is full")
	}

	return nil
}

func (q *queue) Dequeue() *Message {
	q.m_lock.Lock()
	defer q.m_lock.Unlock()

	if len(q.messages) == 0 {
		return nil
	}

	var msg *Message
	select {
	case msg = <-q.messages:
	default:
		return nil
	}

	return msg
}

func (q *queue) AddConsumer(l Link) *Consumer {
	q.c_lock.Lock()
	defer q.c_lock.Unlock()

	c := &Consumer{
		link: l,
	}
	q.consumers = append(q.consumers, c)

	return c
}

func (q *queue) drain() {
	for {
		msg := q.Dequeue()
		if msg == nil {
			return
		}

		q.Deliver(msg)
	}
}

func (q *queue) removeConsumer(i uint64) {
	q.c_lock.Lock()
	defer q.c_lock.Unlock()

	cl := len(q.consumers)
	copy(q.consumers[i:], q.consumers[i+1:])
	q.consumers[cl-1] = nil
	q.consumers = q.consumers[:cl-1]
}
