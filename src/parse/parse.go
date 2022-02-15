package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

type Parser struct {
	Scanner *bufio.Scanner
	Buf     *bytes.Buffer
}

func (pa *Parser) NewParser(r *net.TCPConn) {
	pa.Scanner = bufio.NewScanner(r)
	pa.Buf = &bytes.Buffer{}
}

var HTTPMethods = map[string]bool{
	"POST":    true,
	"GET":     true,
	"HEAD":    true,
	"OPTIONS": true,
}

func isHTTPMethod(text string) bool {
	return HTTPMethods[text]
}

func (pa *Parser) ParseFirstLine() string {
	// var result string

	pa.Scanner.Scan()

	content := strings.Split(pa.Scanner.Text(), " ")
	if !isHTTPMethod(content[0]) {
		log.Printf("%s method is not allowed\n", content[0])
	}
	sendTo := content[1]
	content[1] = "/"
	result := strings.Join(content, " ")
	pa.Buf.WriteString(result)
	pa.Buf.WriteByte('\n')
	return sendTo
}

func (pa *Parser) ParseRest() {
	// var result string
	hostHeaderFound := false

	for pa.Scanner.Scan() {
		content := pa.Scanner.Text()
		if strings.HasPrefix(content, "Host: ") {
			hostHeaderFound = true

		} else if strings.HasPrefix(content, "Proxy-Connection: ") {
			continue
		}
		if content == "" {
			break
		}
		pa.Buf.WriteString(content)
		pa.Buf.WriteByte('\n')
	}
	if !hostHeaderFound {
		log.Printf("No host header found! Aborting!\n")
	}
}

func ParseAddr(raw string) string {
	fmt.Println(raw)
	result := raw[7:]
	var port string

	if raw[:7] == "http://" {
		port = ":80"
	} else if raw[:8] == "https://" {
		port = ":443"
	}
	portSymbolIndex := strings.Index(result, ":")
	if portSymbolIndex != -1 {
		port = ""
	}
	result = strings.TrimSuffix(result, "/") + port

	fmt.Println(result)
	return result
}
