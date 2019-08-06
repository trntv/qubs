package main

import (
	"github.com/sirupsen/logrus"
	"github.com/trntv/qubs/broker"
	"github.com/trntv/qubs/protocol"
	"github.com/trntv/qubs/protocol/grpc"
	"github.com/trntv/qubs/protocol/http"
	"github.com/trntv/qubs/transport"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app   = kingpin.New("Queues hub", "Message broker").Version("0.0.1")
	addr  = app.Flag("addr", "Address to listen").Default("127.0.0.1").Envar("QUBS_ADDR").IP()
	port  = app.Flag("port", "Port to listen").Default("7171").Envar("QUBS_PORT").Uint16()
	proto = app.Flag("proto", "Transport protocol").Default("grpc").Envar("QUBS_PROTOCOL").Enum("grpc", "http")
	debug = app.Flag("debug", "Debug mode").Short('d').Envar("QUBS_DEBUG").Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	b := broker.NewBroker()

	logrus.Infof("Starting %s server at %s:%d", *proto, addr.String(), *port)
	var srv protocol.Server
	switch *proto {
	case "grpc":
		srv = grpc.NewgRPCServer(b)
	case "http":
		srv = http.NewHTTPServer(b)
	}

	err := transport.Listen(addr.String(), *port, srv)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Goodbye!")
}
