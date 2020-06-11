package network

import (
	"errors"
	"net"
	"sync/atomic"
	"time"
)

var (
	ErrWriteOnCloseConn     = errors.New("写入已经关闭的连接")
	ErrWriteTimeOut         = errors.New("连接写超时")
	ErrCloseClosedConn      = errors.New("关闭已经关闭的连接")
	ErrConnectConnectedConn = errors.New("连接已经连接的连接")
	ErrReadNetComplete      = errors.New("接受数据未完成")
	ErrMessageOverBuf       = errors.New("消息超出缓冲限制")
	ErrReadCloseConn        = errors.New("从关闭的连接读取消息")
	ErrConnAlreadyConnected = errors.New("重复建立连接")
)

const (
	MaxPackSize = 64 * 1024
)

//连接唯一ID
var connSequenceID uint64

func Init() {
	connSequenceID = 1
}

func nextSequenceID() uint64 {
	return atomic.AddUint64(&connSequenceID, 1)
}

type Conn interface {
	UniqueId() uint64
	Read() ([]byte, error)
	Write(b []byte) (int, error)
	Close() error
}

type connectedConn interface {
	Conn
	Connect(addr string) error
}

type socket interface {
	String() string       //主要是输出日志用
	RemoteAddr() net.Addr //远端地址
	LocalAddr() net.Addr  //本地地址
	SetFilter(f StreamFilter)

	Close() error //关闭连接

	SetWriteDeadline(t time.Duration) error //写超时设置

	Read(b []byte) (n int, err error)  //读取数据
	Write(b []byte) (n int, err error) //写入数据
}

type StreamFilter interface {
	Read(c net.Conn, b []byte) (int, error)
	Write(c net.Conn, b []byte) (int, error)
}
