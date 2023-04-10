package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// 检查IP地址是否为IPv4地址，并且不是回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// 转为ipv4且是当前首选的ip地址
			if ipnet.IP.To4() != nil && ipnet.IP.IsGlobalUnicast() {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
