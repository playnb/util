package util

import (
	"fmt"
	"github.com/playnb/util/log"
	"github.com/playnb/util/network"
	"testing"
	"time"
)

func init() {
}

func Test_Network_TcpServer(t *testing.T) {
	server := network.NewTcpServer("TestServer", network.DefaultTcpConnOptions)
	server.OnClientConnect = func(conn network.Conn) {
		conn.Write([]byte("hello!"))
		for {
			data, err := conn.Read()
			if err != nil {
				fmt.Println(conn, " read error!", err.Error())
				return
			}
			fmt.Println(conn, " say:", string(data))
			if string(data) == "close" {
				conn.Close()
			}
		}
	}
	server.OnClientLoseConnect = func(conn network.Conn) {
		fmt.Println(conn, " lost")
	}
	server.Listen("127.0.0.1:19000")
	server.WaitAllConnection()
}

func Test_Network_TcpClient(t *testing.T) {
	client := &network.TcpClient{}
	client.Connect("127.0.0.1:19000", network.DefaultTcpConnOptions)
	defer client.Close()

	data, err := client.Conn().Read()
	if err != nil {
		log.Error("%s Read %s", client, err.Error())
		return
	}
	log.Trace("%s <- %s", client, string(data))
	client.Conn().Write([]byte("AAAAA!"))
	client.Conn().Write([]byte("BBBBB!"))
	time.Sleep(time.Second * 3)
	client.Close()
	client.Conn().Write([]byte("CCCCC!"))
}
