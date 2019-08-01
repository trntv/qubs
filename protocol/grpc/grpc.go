package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/trntv/qubs/broker"
	"google.golang.org/grpc"
	"net"
)

func NewgRPCServer(b *broker.Broker) func(net.Listener) {
	return func(listener net.Listener) {

		grpcServer := grpc.NewServer()
		RegisterQueueServer(grpcServer, &queueServer{b: b})

		err := grpcServer.Serve(listener)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}
