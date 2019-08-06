Simple message broker build for development purposes on top of gRPC. Has gRPC and HTTP/2 endpoints.

Run gRPC example:
```
go run cmd/cmd.go --debug --proto=grpc
go run examples/grpc/consumer.go
go run examples/grpc/producer.go
```

Run HTTP example:
```
go run cmd/cmd.go --debug --proto=http --port=80
```
```
curl -X POST \
  http://127.0.0.1:80/some-queue-name \
  -d '{
	"payload": "VGVzdCBQYXlsb2Fk"
}'
```
```
curl http://127.0.0.1:80/some-queue-name
```

TODO:
- [ ] docker image
- [ ] auth
- [ ] messages metadata
- [ ] ack/noack
- [ ] client (consumer, producer)
- [ ] embedded broker
- [ ] usage
- [ ] tests 
- [X] ... a lot of other items already done ...
