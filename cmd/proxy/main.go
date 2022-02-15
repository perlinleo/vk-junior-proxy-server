package main

import (
	"net"

	proxy "github.com/perlinleo/vk-junior-proxy-server/src/proxy"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		panic(err)
	}

	pending, complete := make(chan *net.TCPConn), make(chan *net.TCPConn)

	for i := 0; i < 5; i++ {
		go proxy.HandleConn(pending, complete)
	}

	go proxy.CloseConn(complete)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		pending <- conn
	}
}
