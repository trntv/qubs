package http

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/trntv/qubs/broker"
	"github.com/trntv/qubs/protocol"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"strings"
)

type httpServer struct {
	b *broker.Broker
}

func (s *httpServer) listen(listener net.Listener) {
	srv := http.Server{}
	srv2 := http2.Server{}

	_ = http2.ConfigureServer(&srv, &srv2)
	http.HandleFunc("/", s.handle)
	err := srv.Serve(listener)
	if err != nil {
		logrus.Fatal(err)
	}
}

func NewHTTPServer(b *broker.Broker) protocol.Server {
	s := &httpServer{b: b}
	return s.listen

}

func (s *httpServer) handle(writer http.ResponseWriter, request *http.Request) {
	var hub, queue string
	parts := strings.Split("/", request.URL.Path)
	if len(parts) == 2 {
		hub = parts[0]
		queue = parts[1]
	} else if len(parts) == 1 {
		hub = "default"
		queue = parts[0]
	} else {
		writer.WriteHeader(400)
		return
	}

	switch request.Method {
	case "GET":
		tags := request.URL.Query().Get("tags")
		msg, ok := s.dequeue(hub, queue, strings.Split(tags, ","))
		if !ok {
			writer.WriteHeader(204)
			return
		}

		err := json.NewEncoder(writer).Encode(msg)
		if err != nil {
			logrus.Error(err)
			return
		}
	case "POST":
		msg := &httpMessage{}
		err := json.NewDecoder(request.Body).Decode(msg)
		if err != nil {
			writer.WriteHeader(400)
			_, err := writer.Write([]byte(err.Error()))
			logrus.Error(err)
			return
		}

		err = s.enqueue(hub, queue, msg)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
		}
		writer.WriteHeader(200)
	case "HEAD":
		writer.WriteHeader(200)
	default:
		writer.WriteHeader(501)
	}
}

func (s *httpServer) enqueue(hub string, queue string, msg *httpMessage) error {
	return s.b.GetQueue(hub, queue).Enqueue(broker.NewMessage(msg.Tags, msg.Payload, msg.Broadcast))
}

func (s *httpServer) dequeue(hub string, queue string, tags []string) (*httpMessage, bool) {
	q := s.b.GetQueue(hub, queue)
	msg := q.Dequeue(tags)
	if msg == nil {
		return nil, false
	}

	return &httpMessage{
		Payload: msg.Payload,
	}, true
}
