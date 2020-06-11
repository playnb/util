package network

import (
	"context"
	"fmt"
	"github.com/playnb/util/log"
	"net"
	"time"
)

type tcpConnImpl struct {
	conn         net.Conn
	writeTimeOut time.Duration
	logName      string

	filter StreamFilter
}

func (ti *tcpConnImpl) SetFilter(f StreamFilter) {
	ti.filter = f
}

func (ti *tcpConnImpl) String() string {
	return ti.logName
}

//远端地址
func (ti *tcpConnImpl) RemoteAddr() net.Addr {
	return ti.conn.RemoteAddr()
}

//本地地址
func (ti *tcpConnImpl) LocalAddr() net.Addr {
	return ti.conn.LocalAddr()
}

//关闭连接
func (ti *tcpConnImpl) Close() error {
	return ti.conn.Close()
}

//写超时设置
func (ti *tcpConnImpl) SetWriteDeadline(t time.Duration) error {
	ti.writeTimeOut = t
	return nil
}

//读取数据
func (ti *tcpConnImpl) Read(b []byte) (n int, err error) {
	if ti.filter != nil {
		return ti.filter.Read(ti.conn, b)
	} else {
		return ti.conn.Read(b)
	}
}

//写入数据
func (ti *tcpConnImpl) Write(b []byte) (n int, err error) {
	if ti.filter != nil {
		return ti.filter.Write(ti.conn, b)
	} else {
		return ti.conn.Write(b)
	}
}

type TcpConnOptions struct {
	PendingNum       int
	MaxConnectionNum int
}

func newTCPConn(opts *TcpConnOptions, onClose func(Conn)) *TcpConn {
	c := &TcpConn{}
	c.init(opts)
	c.OnClose = onClose
	return c
}

// 实现TCP连接
type TcpConn struct {
	impl       *connImpl
	tcpImpl    *tcpConnImpl
	lAddr      *net.TCPAddr
	rAddr      *net.TCPAddr
	rAddrRaw   string
	ctx        context.Context
	opts       *TcpConnOptions
	serverSide bool
	logName    string

	OnClose func(Conn)
}

func (tcp *TcpConn) init(opts *TcpConnOptions) {
	tcp.opts = opts
	tcp.logName = "[TcpConn]"

	tcp.tcpImpl = &tcpConnImpl{}
	tcp.tcpImpl.SetFilter(&StreamFilterTcpPack{LittleEndian: false})

	tcp.impl = &connImpl{}
	tcp.impl.Init(tcp.tcpImpl, tcp.ctx, 1024)

	tcp.impl.OnClose = func() {
		if tcp.OnClose != nil {
			tcp.OnClose(tcp)
		}
	}
}

func (tcp *TcpConn) String() string {
	return tcp.logName
}

func (tcp *TcpConn) Connect(addr string) error {
	if tcp.tcpImpl.conn != nil {
		return ErrConnAlreadyConnected
	}
	tcp.rAddrRaw = addr
	tcp.serverSide = false
	err := tcp.connect()
	if err == nil {
		tcp.logName = tcp.impl.String()
	}
	return err
}
func (tcp *TcpConn) ConnectBySocket(socket net.Conn) error {
	if tcp.tcpImpl.conn != nil {
		return ErrConnAlreadyConnected
	}
	tcp.serverSide = true
	err := tcp.connectBySocket(socket)
	if err == nil {
		tcp.logName = tcp.impl.String()
	}
	return err
}

func (tcp *TcpConn) UniqueId() uint64 {
	return tcp.impl.uniqueId
}

func (tcp *TcpConn) Read() ([]byte, error) {
	d, ok := <-tcp.impl.readChan
	if ok {
		return d, nil
	} else {
		return nil, ErrReadCloseConn
	}
}

func (tcp *TcpConn) Write(b []byte) (int, error) {
	return tcp.impl.Write(b)
}

func (tcp *TcpConn) Close() error {
	return tcp.impl.Close()
}

func (tcp *TcpConn) connect() error {
	var err error
	if tcp.tcpImpl.conn != nil {
		return ErrConnectConnectedConn
	}
	tcp.rAddr, err = net.ResolveTCPAddr("", tcp.rAddrRaw)
	if err != nil {
		log.Error("%s connect error:%s", tcp, err.Error())
		return err
	}
	tcp.tcpImpl.conn, err = net.DialTCP("tcp", tcp.lAddr, tcp.rAddr)
	tcp.tcpImpl.logName = fmt.Sprintf("[Tcp@%d=>%s]", tcp.UniqueId(), tcp.rAddrRaw)
	if err != nil {
		log.Error("%s connect error:%s", tcp, err.Error())
		return err
	}
	go tcp.impl.SendLoop()
	go tcp.impl.ReadLoop()
	return nil
}
func (tcp *TcpConn) connectBySocket(socket net.Conn) error {
	tcp.tcpImpl.conn = socket
	tcp.tcpImpl.logName = fmt.Sprintf("[Tcp@%d]", tcp.UniqueId())
	go tcp.impl.SendLoop()
	go tcp.impl.ReadLoop()
	return nil
}
