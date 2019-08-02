package broker

import (
	"errors"
	"github.com/sirupsen/logrus"
	"sort"
	"sync"
	"sync/atomic"
)

const QUEUE_SIZE = 65536 // @todo: move to config

type Queue struct {
	Consumers []*Consumer

	name     string
	Messages chan *Message
	index    uint64
	hub      *Hub
	size     uint64
	lock     sync.Mutex
}

func NewQueue(name string) *Queue {
	q := &Queue{
		name:      name,
		Messages:  make(chan *Message, QUEUE_SIZE),
		Consumers: make([]*Consumer, 0),
	}
	go q.listen()

	return q
}

func (q *Queue) RegisterConsumer(l Link, tags []string) *Consumer {
	c := &Consumer{
		link: l,
		tags: tags,
	}
	q.Consumers = append(q.Consumers, c)

	return c
}

func (q *Queue) Enqueue(msg *Message) error {
	// @todo overflow -> rearrange Messages
	msg.index = atomic.AddUint64(&q.index, 1)
	select {
	case q.Messages <- msg:
		atomic.AddUint64(&q.size, 1)
		return nil
	default:
		return errors.New("queue is full")
	}
}

func (q *Queue) listen() {
	for msg := range q.Messages {
		// @todo credits, block when no consumers
		q.dispatch(msg)
	}
}

func (q *Queue) dispatch(msg *Message) {
	var sent bool
	concopy := make([]*Consumer, 0)
	for _, c := range q.Consumers {
		if c.link.IsClosed() {
			continue
		}

		if len(msg.tags) > 0 && len(c.tags) > 0 {
			var hasTag bool
			for _, t := range msg.tags {
				if c.hasTag(t) {
					hasTag = true
					break
				}
			}

			if !hasTag {
				concopy = append(concopy, c)
				continue
			}
		}

		err := c.link.Send(*msg)
		if err != nil {
			logrus.Error(err)
			continue
		}

		sent = true
		atomic.AddUint64(&c.sent, 1)
		concopy = append(concopy, c)
		if !msg.broadcasting {
			break
		}
	}

	if !sent {
		go q.reqeue(msg)
		return
	} else if len(concopy) > 1 && !msg.broadcasting {
		// sort consumers by number of sent messages (round-robin)
		sort.Slice(concopy, func(i, j int) bool {
			return concopy[i].sent > concopy[j].sent
		})
	}

	q.lock.Lock()
	q.Consumers = concopy
	q.lock.Unlock()

	atomic.AddUint64(&q.size, ^uint64(0))
}

func (q *Queue) reqeue(msg *Message) {
	q.Messages <- msg
}

func (q *Queue) Dequeue(strings []string) *Message {
	// @todo dequeue
	return nil
}
