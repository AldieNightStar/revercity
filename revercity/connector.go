package revercity

import "net"

type Connector interface {
	Connect() (net.Conn, error)
}

// ===================
// TCP Connector
// ===================

type TCPConnector struct {
	address string
}

func (c *TCPConnector) Connect() (net.Conn, error) {
	return net.Dial("tcp", c.address)
}

func NewTcpConnector(addr string) *TCPConnector {
	return &TCPConnector{address: addr}
}
