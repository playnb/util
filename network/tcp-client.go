package network

import "fmt"

type TcpClient struct {
	conn    connectedConn
	addr    string
	opts    *TcpConnOptions
	logName string

	OnLoseConnection func()
}

func (c *TcpClient) String() string {
	return c.logName
}

func (c *TcpClient) Connect(addr string, opts *TcpConnOptions) error {
	c.logName = fmt.Sprintf("[Client=>%s]", c.addr)
	c.opts = opts
	c.conn = newTCPConn(c.opts, func(Conn) {
		if c.OnLoseConnection != nil {
			c.OnLoseConnection()
		}
	})
	return c.conn.Connect(addr)
}

func (c *TcpClient) Conn() Conn {
	return c.conn
}

func (c *TcpClient) UniqueId() uint64 {
	return c.conn.UniqueId()
}

func (c *TcpClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *TcpClient) Write(b []byte) (int, error) {
	return c.conn.Write(b)
}

func (c *TcpClient) Read() ([]byte, error) {
	return c.conn.Read()
}
