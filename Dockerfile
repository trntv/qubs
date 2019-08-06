# Build binary
FROM golang:1.12-alpine AS build
RUN apk add --no-cache git
WORKDIR /go/src/github.com/trntv/qubs
ADD . .
ENV GO111MODULE=on
RUN go build -o qubs

# Build image
FROM alpine
WORKDIR /app
COPY --from=build /go/src/github.com/trntv/qubs/qubs /app/
ENTRYPOINT ./qubs
