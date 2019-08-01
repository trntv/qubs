package grpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trntv/qubs/broker"
	"google.golang.org/grpc/metadata"
	"io"
)

type queueServer struct {
	b *broker.Broker
}

func (qs *queueServer) Enqueue(ctx context.Context, msg *Message) (*Ack, error) {
	hubs := resolveHubs(ctx)
	for _, hub := range hubs {
		cmsg := broker.NewMessage(msg.Tags, msg.Payload, false)
		q := qs.b.GetQueue(hub, msg.Queue)
		_ = q.Enqueue(cmsg) // @todo: errors handling
	}

	return &Ack{Id: msg.Id}, nil
}

func (qs *queueServer) EnqueueBatch(stream Queue_EnqueueBatchServer) error {
	hubs := resolveHubs(stream.Context())
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("%v\r\n", msg)

		for _, hub := range hubs {
			cmsg := broker.NewMessage(msg.Tags, msg.Payload, false)
			q := qs.b.GetQueue(hub, msg.Queue)
			_ = q.Enqueue(cmsg) // @todo: errors handling
		}
	}

	return nil
}

func (qs *queueServer) Consume(msg *ConsumerConnect, stream Queue_ConsumeServer) error {
	l := &consumerLink{
		stream: stream,
	}
	hubs := resolveHubs(stream.Context())

	for _, hub := range hubs {
		q := qs.b.GetQueue(hub, msg.Queue)
		q.RegisterConsumer(l, msg.Tags)
	}

	err := l.run()
	logrus.Error(err)

	return nil
}

type consumerLink struct {
	stream Queue_ConsumeServer
	q chan bool
	closed bool
}

func (l *consumerLink) Send(msg broker.Message) error {
	return l.stream.Send(&Message{
		Payload: msg.Payload,
	})
}

func (l *consumerLink) Close() error {
	return nil
}

func (l *consumerLink) IsClosed() bool {
	return l.closed
}

func (l *consumerLink) run() error {
	m := new(ConsumerConnect)
	for {
		if err := l.stream.RecvMsg(m); err != nil {
			if err != io.EOF {
				l.closed = true
				return err
			}
		}
	}

	return nil
}

func resolveHubs(ctx context.Context) []string {
	hubs := []string{"default"}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return hubs
	}

	mhubs := md.Get("hubs")
	if len(mhubs) > 0 {
		hubs = mhubs
	}

	return hubs
}


