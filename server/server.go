package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	go RunBroadcastMatchClient()
	println("input client ip:")
	in := bufio.NewReader(os.Stdin)
	inputIp, _ := in.ReadString('\n')
	// ReadString读\n结束并接收\n，此处去除最后的\n,windows是\r\n
	inputIp = strings.TrimSuffix(inputIp, "\n")
	inputIp = strings.TrimSuffix(inputIp, "\r")
	conn, err := net.Dial("udp", inputIp) // 目标IP地址和端口号
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer conn.Close()

	time.Sleep(time.Second * 5)
	msg := []byte("space from sender!")
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Message sent.")

}
