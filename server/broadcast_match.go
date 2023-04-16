package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var c rune

func RunBroadcastMatchClient() {
	/*
		// 服务端若存在多个网络适配器（如虚拟适配Vmware网卡、多个有效本地连接），
		// 则会导致无法接收客户端的配对，因为有可能本机IP优先级被其他适配器给“盗走”了。
		// 反过来也是如此。此问题无解，因为没有独立出一个服务端做中转，且此项目需求也不需要独立出服务端
		// 因此考虑让服务端增加一个输入框，指定输入客户端的IP是此情景的一个常见的选项。
	*/
	startListenPacket(net.IPv4zero.String() + ":1234")
}

func startListenPacket(ip string) {

	conn, err := net.ListenPacket("udp", ip)
	if err != nil {
		println("致命错误：net.ListenPacket\n服务端接收客户端配对信息监听失败" + err.Error())
		fmt.Scanln(&c)
		panic("")
	}

	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		err = conn.SetReadDeadline(time.Now().Add(6 * time.Second))
		if err != nil {
			println("致命错误：SetReadDeadline\n服务端接收客户端配对信息的超时设置失败:" + err.Error())
			fmt.Scanln(&c)
			panic("")
		}
		// 接收服务端广播发送的具体消息
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// 超时，重新等待下一轮广播消息
				println("超时，重新等待下一轮广播消息")
				continue
			}
			println("致命错误：conn.ReadFrom\n服务端接收客户端配对信息读取失败:" + err.Error())
			fmt.Scanln(&c)
			panic("")
		}
		ms := string(buf[:n])
		if !strings.Contains(ms, "key_sprite_match") {
			continue
		}
		name := strings.Split(ms, ",")[1]
		fmt.Printf("Received client broadcast message from ip:%s\n"+
			"name:%s\nmsg:%s\n", addr.String(), name, ms)
	}
}
