// +build examples

package main

import (
	"context"
	"fmt"
	g "github.com/trntv/qubs/protocol/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:7171", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := g.NewQueueClient(conn)

	ctx := context.Background()
	md := metadata.AppendToOutgoingContext(ctx, "host", "hub1")

	_, err = client.Enqueue(md, &g.Message{Queue: "q1", Payload: []byte("Message for q1")})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message 1 sent")

	_, err = client.Enqueue(md, &g.Message{Queue: "q2", Payload: []byte("Message for q2")})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message 2 sent")
}
