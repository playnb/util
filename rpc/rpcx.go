package rpc

import (
	"github.com/playnb/util/log"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

func NewRpcxServ(addr string, updateInterval time.Duration, etcdService []string, etcdBase string) *server.Server {
	var err error
	s := server.NewServer()

	if len(etcdService) > 0 {
		r := &serverplugin.EtcdV3RegisterPlugin{
			ServiceAddress: "tcp@" + addr,
			EtcdServers:    etcdService,
			BasePath:       etcdBase,
			Metrics:        metrics.NewRegistry(),
			UpdateInterval: updateInterval,
		}
		err = r.Start()
		if err != nil {
			log.Error("%s 初始化Etcd插件 %s", s, err.Error())
			panic(err)
		}
		s.Plugins.Add(r)
	}

	go func() {
		err = s.Serve("tcp", addr)
		if err != nil && err != server.ErrServerClosed {
			log.Error("%s 初始化失败 %s", s, err.Error())
			panic(err)
		}
	}()

	return s
}
