package main

import (
	"fmt"
	"key_sprite/common"
	"net"
	"time"
)

var c rune

func RunClientBroadcast() {
	//localIp := common.GetLocalActiveIPs()[0]
	//localBroadcastIp := common.ResolveLocalBroadcastIp(localIp)
	//port := "1234"
	//localBroadcastIpAndPort := localBroadcastIp + ":" + port
	// 不间断发送配对信息的广播
	for true {
		time.Sleep(time.Second * 5)
		addr := "255.255.255.255:1234" // 广播地址和端口号
		// 固定端口而不是随机
		localAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 1234}
		udpAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			fmt.Println("Error resolving UDP address:", err.Error())
			return
		}
		// DialUDP取代Dial解决端口自动随机问题
		conn, err := net.DialUDP("udp", localAddr, udpAddr)
		if err != nil {
			println("致命错误：RunClientBroadcast\n不间断发送配对信息的广播, net.Dial连接错误:" + err.Error())
			fmt.Scanln(&c)
			panic("")
		}

		msg := []byte("key_sprite_match," + common.LocalComputerName())
		_, err = conn.Write(msg)
		if err != nil {
			println("致命错误：RunClientBroadcast\n不间断发送配对信息的广播, conn.Write发送错误:" + err.Error())
			fmt.Scanln(&c)
			panic("")
		}
		fmt.Println("match Message sent.")
		func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
			}
		}(conn)
	}
	fmt.Scanln(&c)
}
