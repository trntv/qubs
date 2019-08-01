package transport

import (
	"fmt"
	"github.com/trntv/qubs/protocol"
	"net"
)

func Listen(addr string, port uint16, handler protocol.Server) error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}
	defer ln.Close()

	handler(ln)

	return nil
}
