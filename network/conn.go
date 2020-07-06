package network

import (
	"context"
	"github.com/playnb/util/log"
	"net"
	"sync"
	"time"
)

/*
主要处理连接和收发消息
*/

type connImpl struct {
	writeChan chan []byte //发送缓冲
	readChan  chan []byte //接受缓冲

	connMutex    sync.Mutex
	closeFlag    bool //关闭标识
	uniqueId     uint64
	writeTimeOut time.Duration

	socket  socket
	ctx     context.Context
	sending bool
	reading bool
	inited  bool

	OnClose func()
}

func (c *connImpl) Init(socket socket, ctx context.Context, pendingNum int) {
	c.uniqueId = nextSequenceID()
	c.closeFlag = false
	c.writeChan = make(chan []byte, pendingNum)
	c.readChan = make(chan []byte, pendingNum)
	c.ctx = ctx
	c.socket = socket

	c.socket.SetWriteDeadline(time.Second * 30)
	c.writeTimeOut = 3 * time.Second
	c.inited = true
}

func (c *connImpl) SendLoop() {
	defer func() {
		log.Debug("%s SendLoop 结束, 关闭连接", c)
		c.sending = false
		c.Close()
	}()
	c.sending = true
	log.Trace("%s 开始SendLoop", c)
	for {
		select {
		case b, ok := <-c.writeChan:
			if ok == false {
				log.Debug("%s SendLoop关闭发送channel writeChan", c)
				return
			}
			if b == nil {
				log.Debug("%s nil消息 主动结束发送数据goroutine", c)
				//return
				break
			}
			if c.closeFlag == true {
				log.Error("%s SendLoop Error:向关闭的端口写数据", c)
				//return
				break
			}

			len, err := c.socket.Write(b)
			if err != nil {
				if len > 0 {
					//虽然错误但是也发送了部分数据
				}
				log.Error("%s SendLoop发送数据错误 %s", c, err.Error())
				//return
				break
			}
		}
	}
}
func (c *connImpl) ReadLoop() {
	defer func() {
		log.Debug("%s ReadLoop 结束, 关闭连接", c)
		c.reading = false
		close(c.readChan)
		c.Close()
	}()
	c.reading = true
	readBuf := []byte(nil)
	for {
		if readBuf == nil {
			readBuf = make([]byte, MaxPackSize)
		}
		n, err := c.socket.Read(readBuf)
		if err == ErrReadNetComplete {
			continue
		}
		if err != nil {
			log.Error("%s ReadLoop Err:%s", c, err.Error())
			return
		}
		if n == 0 {
			//读到0数据
			log.Error("%s ReadLoop 读到0数据", c)
			return
		}
		c.readChan <- readBuf[:n]
		readBuf = nil
	}
}

func (c *connImpl) RemoteAddr() net.Addr { //远端地址
	return c.socket.RemoteAddr()
}
func (c *connImpl) LocalAddr() net.Addr { //本地地址
	return c.socket.LocalAddr()
}
func (c *connImpl) String() string {
	return c.socket.String()
}
func (c *connImpl) Read(b []byte) (n int, err error) {
	return c.socket.Read(b)
}
func (c *connImpl) Write(b []byte) (n int, err error) {
	//c.connMutex.Lock()
	//defer c.connMutex.Unlock()
	if c.closeFlag == true {
		log.Error("[%s] TCPConn:Write 向关闭的Conn写数据", c)
		return 0, ErrWriteOnCloseConn
	}

	if c.writeTimeOut > 0 {
		timeOut := time.NewTimer(c.writeTimeOut)
		select {
		case c.writeChan <- b:
			return len(b), nil
		case <-timeOut.C:
			log.Error("%s Write 发送消息超时", c)
			return 0, ErrWriteOnCloseConn
		}
	} else {
		c.writeChan <- b
		return len(b), nil
	}
}
func (c *connImpl) Close() error {
	c.connMutex.Lock()
	defer c.connMutex.Unlock()
	if c.closeFlag {
		if !c.reading && !c.sending {
			log.Error("%s Close 已经关闭的连接", c)
		}
		return ErrCloseClosedConn
	}
	if c.reading && c.sending {
		log.Trace("%s 主动关闭", c)
	} else {
		log.Trace("%s 被动关闭 reading:%v sending:%v", c, c.reading, c.sending)
	}
	c.closeFlag = true
	c.socket.Close()
	if c.OnClose != nil {
		c.OnClose()
	}
	return nil
}
func (c *connImpl) SetWriteDeadline(t time.Duration) error { //写超时设置
	c.writeTimeOut = t
	c.socket.SetWriteDeadline(t * 10)
	return nil
}
