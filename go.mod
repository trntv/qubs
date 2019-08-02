module github.com/trntv/qubs

go 1.12

replace qpid.apache.org/amqp => ./vendor/qpid.apache.org/amqp

replace qpid.apache.org/electron => ./vendor/qpid.apache.org/electron

replace qpid.apache.org/proton => ./vendor/qpid.apache.org/proton

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.22.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	qpid.apache.org/amqp v0.0.0-00010101000000-000000000000
	qpid.apache.org/electron v0.0.0-00010101000000-000000000000
	qpid.apache.org/proton v0.0.0-00010101000000-000000000000
)
