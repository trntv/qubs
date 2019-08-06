Simple messages broker built for DEVELOPMENT purposes on top of gRPC. Has gRPC and HTTP/2 endpoints.

Run gRPC example:
```
go run cmd/cmd.go --debug --proto=grpc
```
or
```
docker run --rm -p 7171:7171 trntv/qubs:latest --proto=grpc
```

```
go run examples/grpc/consumer.go
go run examples/grpc/producer.go
```


Run HTTP example:
```
go run cmd/cmd.go --debug --proto=http --port=8181
```
or
```
docker run --rm -p 8181:8181 trntv/qubs:latest --proto=http --port=8181
```
```
curl -X POST \
  http://127.0.0.1:8181/some-queue-name \
  -d '{
	"payload": "VGVzdCBQYXlsb2Fk"
}'
```
```
curl http://127.0.0.1:8181/some-queue-name
```

TODO:
- [X] docker image
- [ ] auth
- [ ] messages metadata
- [ ] ack/noack
- [ ] client (consumer, producer)
- [ ] embedded broker
- [ ] usage
- [ ] tests 
- [X] ... a lot of other items already done ...
