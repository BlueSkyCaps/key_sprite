package common

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// GetLocalActiveIPs 获取本机可用ip地址列表
func GetLocalActiveIPs() []string {
	addrsTmp, _ := net.InterfaceAddrs()
	var addrs []string
	for _, address := range addrsTmp {
		// 检查IP地址是否为IPv4地址，并且不是回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// 转为ipv4且是当前首选的ip地址
			if ipnet.IP.To4() != nil && ipnet.IP.IsGlobalUnicast() {
				addrs = append(addrs, ipnet.IP.String())
				fmt.Println(ipnet.IP.String())
			}
		}
	}
	if len(addrs) < 1 {
		panic("致命错误：GetLocalActiveIPs\n没有获取到任何本机IP")
	}
	return addrs
}

func ResolveLocalBroadcastIp(localActiveIp string) string {
	localIpSegment := strings.Split(localActiveIp, ".")[:3]
	localBroadcastIp := strings.Join(localIpSegment, ".") + ".255"
	return localBroadcastIp
}

// GetNetActiveIPs 根据本机网络设备获取ip地址
func GetNetActiveIPs() []net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	var ips []net.IP
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				continue
			}

			ips = append(ips, ip)
		}
	}
	return ips
}

func LocalComputerName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown device name"
	}
	return hostname
}
