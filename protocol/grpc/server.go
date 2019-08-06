package grpc

import (
	"context"
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
	cmsg := broker.NewMessage(msg.Payload, false)

	for _, hub := range hubs {
		q := qs.b.GetQueue(hub, msg.Queue)
		_ = q.Enqueue(cmsg)
	}

	return &Ack{Seq: cmsg.Id()}, nil
}

func (qs *queueServer) Consume(msg *ConsumerConnect, stream Queue_ConsumeServer) error {
	l := &consumerLink{
		stream: stream,
	}
	hubs := resolveHubs(stream.Context())

	for _, hub := range hubs {
		q := qs.b.GetQueue(hub, msg.Queue)
		q.AddConsumer(l)
	}

	err := l.run()
	logrus.Error(err)

	return nil
}

type consumerLink struct {
	stream Queue_ConsumeServer
	q      chan bool
	closed bool
}

func (l *consumerLink) Send(msg *broker.Message) error {
	return l.stream.Send(&Message{
		Payload: msg.Payload(),
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
