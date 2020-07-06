package util

import (
	"github.com/playnb/util/log"
	"net"
)

func GetIp(ifName string) string {
	interfaces, err := net.Interfaces()
	defaultIp := "127.0.0.1"
	if err != nil {
		return defaultIp
	}

	for _, i := range interfaces {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			log.Trace("[Util] GetIp(%s) Error:%s", i.Name, err.Error())
			continue
		}
		addresses, err := byName.Addrs()
		for _, v := range addresses {
			if ipNet, ok := v.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					defaultIp = ipNet.IP.String()
					if byName.Name == ifName {
						return ipNet.IP.String()
					}
				}
			}
		}
	}
	return defaultIp
}

func ShowAllIp() {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, i := range interfaces {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			continue
		}
		addresses, err := byName.Addrs()
		for _, v := range addresses {
			if ipNet, ok := v.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					log.Trace("[Util] ShowAllIp (%s)  (%s)", byName.Name, ipNet.IP.To4().String())
				}
			}
		}
	}
}
