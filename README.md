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
- [X] hubs
- [x] tags
- [x] cli command
- [ ] client (consumer, producer)
- [ ] batch publish
- [ ] ack/noack
- [ ] embedded broker
- [ ] usage
- [ ] .....
