package proxy

import (
	"encoding/hex"
	"io"
	"log"
	"net"

	parse "github.com/perlinleo/vk-junior-proxy-server/src/parse"
)

func ProxyConn(conn *net.TCPConn) error {

	var pa parse.Parser
	pa.NewParser(conn)

	sendTo := pa.ParseFirstLine()
	pa.ParseRest()

	rConn, err := net.Dial("tcp", parse.ParseAddr(sendTo))

	if err != nil {
		return err
	}
	defer rConn.Close()
	defer conn.Close()

	// end of request

	pa.Buf.WriteString("\r\n\r\n")

	if _, err := rConn.Write(pa.Buf.Bytes()); err != nil {
		return err
	}
	log.Printf("sent:\n%v", hex.Dump(pa.Buf.Bytes()))

	data := make([]byte, 1024)
	n, err := rConn.Read(data)
	if err != nil {
		if err != io.EOF {
			return err
		} else {
			log.Printf("received err: %v", err)
		}
	}
	log.Printf("received:\n%v", hex.Dump(data[:n]))

	if _, err := conn.Write(data[:n]); err != nil {
		return err
	}
	return nil
}

func HandleConn(in <-chan *net.TCPConn, out chan<- *net.TCPConn) {
	for conn := range in {
		err := ProxyConn(conn)
		if err != nil {
			log.Printf("received err: %v", err)
		}
		out <- conn
	}
}

func CloseConn(in <-chan *net.TCPConn) {
	for conn := range in {
		conn.Close()
	}
}
