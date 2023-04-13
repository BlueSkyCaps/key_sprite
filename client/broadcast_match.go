package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var c rune

func RunBroadcastClientMatch() {
	//localIp := common.GetLocalActiveIPs()[0]
	//localBroadcastIp := common.ResolveLocalBroadcastIp(localIp)
	//port := "1234"
	//localBroadcastIpAndPort := localBroadcastIp + ":" + port
	// 接收广播消息配对服务端的ip
	//actNetIps := common.GetNetActiveIPs()
	//for _, actNetIp := range actNetIps {
	//	if actNetIp.To4() != nil {
	//		startIp := actNetIp.String() + ":" + port
	//		go startListenPacket(startIp)
	//	}
	//}
	//udpAddr, _ := net.ResolveUDPAddr("udp", net.IPv4zero .String()+":1234")
	go startListenPacket(net.IPv4zero.String() + ":1234")

}

func startListenPacket(ip string) {

	conn, err := net.ListenPacket("udp", ip)
	if err != nil {
		println("致命错误：net.ListenPacket\n客户端接收服务端配对信息监听失败" + err.Error())
		fmt.Scanln(&c)
		panic("")
	}

	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		err = conn.SetReadDeadline(time.Now().Add(6 * time.Second))
		if err != nil {
			println("致命错误：SetReadDeadline\n客户端接收服务端配对信息的超时设置失败:" + err.Error())
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
			println("致命错误：conn.ReadFrom\n客户端接收服务端配对信息读取失败:" + err.Error())
			fmt.Scanln(&c)
			panic("")
		}
		ms := string(buf[:n])
		if !strings.Contains(ms, "key_sprite_match") {
			continue
		}
		name := strings.Split(ms, ",")[1]
		fmt.Printf("Received server broadcast message from ip:%s\n"+
			"name:%s\nmsg:%s\n", addr.String(), name, ms)
	}
}
