package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"net"
	"os"
	"time"
)

const (
	VkSpace          = 0x20
	KeyEventFKeydown = 0x0000
	KeyEventFKeyup   = 0x0002
)

var userDll = windows.NewLazyDLL("user32.dll")

func main() {
	keyProc := userDll.NewProc("keybd_event")
	port := ":1234" // 监听端口
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	err = listener.SetDeadline(time.Time{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	for true {
		handleClient(listener, keyProc)
	}

}

func handleClient(conn *net.UDPConn, proc *windows.LazyProc) {
	buf := make([]byte, 1024)
	// 等待远端的发送，然后读取，此条代码阻塞并且有超时机制
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}
	msg := string(buf[:n])
	println("Received ip from:" + addr.IP.String())
	fmt.Printf("Received msg:%s\n", msg)
	pressKey(VkSpace, proc)
}

func pressKey(keyCode int, proc *windows.LazyProc) {
	proc.Call(uintptr(keyCode), 0, KeyEventFKeydown, 0)
	proc.Call(uintptr(keyCode), 0, KeyEventFKeyup, 0)
}
