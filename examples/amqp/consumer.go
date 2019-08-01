/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"qpid.apache.org/amqp"
	"qpid.apache.org/electron"
)

func main() {

	fatalIf := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	var count = flag.Int("count", 1, "Stop after receiving this many messages in total")
	var prefetch = flag.Int("prefetch", 0, "enable a pre-fetch window for flow control")

	flag.Parse()

	debugf := func(format string, data ...interface{}) { log.Printf(format, data...) }
	urlStr := "amqp://127.0.0.1:7171/test-queue"

	messages := make(chan amqp.Message) // Channel for messages from goroutines to main()
	defer close(messages)

	var wait sync.WaitGroup // Used by main() to wait for all goroutines to end.
	wait.Add(1)             // Wait for one goroutine per URL.

	container := electron.NewContainer(fmt.Sprintf("receive[%v]", os.Getpid()))
	connections := make(chan electron.Connection, 1) // Connections to close on exit

	debugf("Connecting to %s\n", urlStr)
	go func(urlStr string) { // Start the goroutine
		defer wait.Done() // Notify main() when this goroutine is done.
		url, err := amqp.ParseURL(urlStr)
		fatalIf(err)
		c, err := container.Dial("tcp", url.Host) // NOTE: Dial takes just the Host part of the URL
		fatalIf(err)
		connections <- c // Save connection so we can Close() when main() ends
		addr := strings.TrimPrefix(url.Path, "/")
		opts := []electron.LinkOption{electron.Source(addr)}
		if *prefetch > 0 { // Use a pre-fetch window
			opts = append(opts, electron.Capacity(*prefetch), electron.Prefetch(true))
		} else { // Grant credit for all expected messages at once
			opts = append(opts, electron.Capacity(*count), electron.Prefetch(false))
		}
		r, err := c.Receiver(opts...)
		fatalIf(err)
		// Loop receiving messages and sending them to the main() goroutine
		for {
			if rm, err := r.Receive(); err == nil {
				fmt.Println("New message!")
				rm.Accept()
				messages <- rm.Message
			} else if err == electron.Closed {
				return
			} else {
				log.Fatalf("receive error %v: %v", urlStr, err)
			}
		}
	}(urlStr)

	// All goroutines are started, we are receiving messages.
	fmt.Printf("Listening on %d connections\n", 1)

	// print each message until the count is exceeded.
	for i := 0; i < *count; i++ {
		m := <-messages
		debugf("%v\n", m.Body())
	}
	fmt.Printf("Received %d messages\n", *count)

	c := <-connections
	debugf("close %s", c)
	c.Close(nil)
	wait.Wait() // Wait for all goroutines to finish.
}
