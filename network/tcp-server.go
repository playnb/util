package network

import (
	"fmt"
	"github.com/playnb/util/log"
	"net"
	"sync"
)

var DefaultTcpConnOptions = &TcpConnOptions{
	PendingNum:       1024,
	MaxConnectionNum: 0,
}

func NewTcpServer(name string, opts *TcpConnOptions) *TcpServer {
	s := &TcpServer{}
	s.Init(name, opts)
	return s
}

type TcpServer struct {
	addr    string
	name    string
	logName string
	running bool
	ln      net.Listener

	allConn map[uint64]Conn
	mutex   sync.Mutex
	opts    *TcpConnOptions
	group   *sync.WaitGroup

	OnClientConnect     func(conn Conn)
	OnClientLoseConnect func(conn Conn)
}

func (s *TcpServer) Init(name string, opts *TcpConnOptions) {
	s.name = name
	s.opts = opts
	s.logName = fmt.Sprintf("[%s@%s]", s.name, s.addr)
	s.allConn = make(map[uint64]Conn)
	s.group = &sync.WaitGroup{}
}

func (s *TcpServer) String() string {
	return s.logName
}

func (s *TcpServer) Listen(addr string) error {
	var err error
	s.addr = addr
	s.logName = fmt.Sprintf("[%s@%s]", s.name, s.addr)
	s.ln, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Error("%s Listen Err:%s", s, err.Error())
		return err
	}
	log.Trace("%s 绑定服务器端口成功", s)
	s.group.Add(1)
	go s.loop()
	return nil
}

func (s *TcpServer) WaitAllConnection() {
	s.group.Wait()
}

func (s *TcpServer) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.group.Done()
	s.ln.Close()
	for _, conn := range s.allConn {
		conn.Close()
	}
	s.allConn = make(map[uint64]Conn)
}

func (s *TcpServer) onClientLoseConnect(conn Conn) {
	if s.OnClientLoseConnect != nil {
		s.OnClientLoseConnect(conn)
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.allConn, conn.UniqueId())
}

func (s *TcpServer) loop() {
	defer func() {
		log.Trace("%s exit loop", s)
		s.running = false
		s.group.Done()
	}()
	s.running = true
	for {
		sock, err := s.ln.Accept()
		if err != nil {
			log.Error("%s Accept Error:%s", s, err.Error())
			continue
		}
		conn := newTCPConn(s.opts, s.onClientLoseConnect)
		conn.ConnectBySocket(sock)

		func() {
			s.mutex.Lock()
			defer s.mutex.Unlock()
			if s.opts.MaxConnectionNum > 0 && len(s.allConn) > s.opts.MaxConnectionNum {
				log.Error("%s Accept Error:超过最大连接数:%d", s, len(s.allConn))
				return
			}
			if _, ok := s.allConn[conn.UniqueId()]; ok {
				log.Error("%s Accept Error:拦截UniqueId重复:%d", s, conn.UniqueId())
				return
			}
			s.allConn[conn.UniqueId()] = conn
			s.group.Add(1)
			if s.OnClientConnect != nil {
				go s.OnClientConnect(conn)
			}
		}()
	}
}
