package broker

import (
	"sync"
	"sync/atomic"
)

type Queue interface {
	Deliver(msg *Message) error
	Consume(fn func(*Message) error)
}

type queue struct {
	seq uint64

	messages  map[uint64]*Message
	consumers []*Consumer

	head, tail uint64

	m_lock sync.Mutex
	c_lock sync.Mutex
}

func NewQueue() *queue {
	q := &queue{
		messages: make(map[uint64]*Message),
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
	q.m_lock.Lock()
	defer q.m_lock.Unlock()

	head := q.head
	q.head = q.head + 1
	q.messages[head] = msg

	msg.delivery.attempt++
	msg.delivery.index = head

	return nil
}

func (q *queue) Dequeue() *Message {
	q.m_lock.Lock()
	defer q.m_lock.Unlock()

	if q.tail == q.head {
		return nil
	}

	tail := q.tail
	q.tail = q.tail + 1

	return q.messages[tail]
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
