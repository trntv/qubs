package protocol

import "net"

type Server func(listener net.Listener)
