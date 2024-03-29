// +build examples

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
	"qpid.apache.org/amqp"
	"qpid.apache.org/electron"
	"strings"
	"sync"
)

func main() {

	fatalIf := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	var count = flag.Int64("count", 1, "Send this many messages to each address.")
	var debug = flag.Bool("debug", true, "Print detailed debug output")
	var Debugf = func(format string, data ...interface{}) {} // Default no debugging output

	flag.Parse()

	if *debug {
		Debugf = func(format string, data ...interface{}) { log.Printf(format, data...) }
	}

	urlStr := "amqp://127.0.0.1:7171/test-queue"

	sentChan := make(chan electron.Outcome) // Channel to receive acknowledgements.

	var wait sync.WaitGroup
	wait.Add(1) // Wait for one goroutine per URL.

	container := electron.NewContainer(fmt.Sprintf("send[%v]", os.Getpid()))
	connections := make(chan electron.Connection, 1) // Connections to close on exit

	Debugf("Connecting to %v\n", urlStr)
	go func(urlStr string) {
		defer wait.Done() // Notify main() when this goroutine is done.
		url, err := amqp.ParseURL(urlStr)
		fatalIf(err)
		c, err := container.Dial("tcp", url.Host) // NOTE: Dial takes just the Host part of the URL
		fatalIf(err)
		connections <- c // Save connection so we can Close() when main() ends
		addr := strings.TrimPrefix(url.Path, "/")
		c.Connection()
		s, err := c.Sender(electron.Target(addr))
		fatalIf(err)
		// Loop sending messages.
		for i := int64(0); i < *count; i++ {
			m := amqp.NewMessage()
			body := fmt.Sprintf("%v%v", addr, i)
			m.Marshal(body)
			s.SendAsync(m, sentChan, body) // Outcome will be sent to sentChan
		}
	}(urlStr)

	// Wait for all the acknowledgements
	expect := int(*count)
	Debugf("Started senders, expect %v acknowledgements\n", expect)
	for i := 0; i < expect; i++ {
		out := <-sentChan // Outcome of async sends.
		if out.Error != nil {
			log.Fatalf("acknowledgement[%v] %v error: %v", i, out.Value, out.Error)
		} else if out.Status != electron.Accepted {
			log.Fatalf("acknowledgement[%v] unexpected status: %v", i, out.Status)
		} else {
			Debugf("acknowledgement[%v]  %v (%v)\n", i, out.Value, out.Status)
		}
	}
	fmt.Printf("Received all %v acknowledgements\n", expect)

	wait.Wait() // Wait for all goroutines to finish.
	close(connections)
	for c := range connections { // Close all connections
		if c != nil {
			c.Close(nil)
		}
	}
}
