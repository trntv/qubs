package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	g "github.com/trntv/qubs/protocol/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"sync"
)

var client g.QueueClient
var ctx context.Context
var wg sync.WaitGroup

func main() {
	conn, err := grpc.Dial("127.0.0.1:7171", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client = g.NewQueueClient(conn)
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs("host", "hub1"))
	wg = sync.WaitGroup{}
	wg.Add(2)

	go startConsumer("q1")
	go startConsumer("q2")

	wg.Wait()
}

func startConsumer(q string) {
	stream, err := client.Consume(ctx, &g.ConsumerConnect{
		Queue: q,
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				logrus.Fatalln(err)
			}
			continue
		}

		fmt.Println(q, msg)
		break
	}

	stream.CloseSend()
	wg.Done()
}
