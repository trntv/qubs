package main

import (
	"github.com/sirupsen/logrus"
	"github.com/trntv/qubs/broker"
	"github.com/trntv/qubs/protocol"
	"github.com/trntv/qubs/protocol/amqp1"
	"github.com/trntv/qubs/protocol/grpc"
	"github.com/trntv/qubs/transport"
	"gopkg.in/alecthomas/kingpin.v2"
)


var (
	addr = kingpin.Flag("addr", "Address to listen").Default("127.0.0.1").Envar("QUBS_ADDR").IP()
	port = kingpin.Flag("port", "Port to listen").Default("7171").Envar("QUBS_PORT").Uint16()
	proto = kingpin.Flag("proto", "Transport protocol").Default("grpc").Envar("QUBS_PROTOCOL").Enum("amqp", "grpc")
	debug = kingpin.Flag("debug", "Debug mode").Short('d').Envar("QUBS_DEBUG").Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	b := broker.NewBroker()

	var srv protocol.Server
	switch *proto {
	case "grpc":
		logrus.Info("Starting gRPC server")
		srv = grpc.NewgRPCServer(b)
	case "amqp":
		logrus.Info("Starting AMQP server")
		srv = amqp1.NewAmqpServer(b)

	}

	err := transport.Listen(addr.String(), *port, srv)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Goodbye!")
}
