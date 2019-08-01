Run:
```
go run cmd/cmd.go --debug
go run example/consumer.go
go run example/producer.go
```

Features:
- hubs, queues, tags
- gRPC-based queue
- AMQP 1.0
- round-robin, fanout
- batch publishing

TODO:
- [ ] Auth
- [ ] HTTP/2 transport (push + pull)
- [X] hubs
- [x] tags
- [x] cli command
- [ ] batch publish
- [ ] ack/noack
- [ ] client (consumer, producer)
- [ ] embedded broker
- [ ] usage
- [ ] tests 
- [ ] .....
