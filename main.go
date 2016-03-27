package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	gateway = "http://127.0.0.1:8080/ipfs/"
	tracker = "104.131.124.184:9090"
)

func Client(host string) net.Conn {
	servAddr := host + ":80"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatal("ResolveTCPAddr failed:", err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Dial failed:", err.Error())
	}

	return conn
}

func handleRequest(conn net.Conn) {
	bufConn := bufio.NewReader(conn)

	req := newRequest()
	err := req.Decode(bufConn)
	if err != nil {
		return
	}

	hash, ok := getFromTracker(req.Header["Host"] + req.Path)
	if ok {
		resp, err := http.Get(gateway + hash.Hash)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		_, err = io.Copy(conn, resp.Body)
		if err != nil {
			log.Fatal("cdn: ", err.Error())
		}

		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!") // Just to show that we proxied something
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	} else {
		client := Client(req.Header["Host"])
		client.Write(req.toBytes())
		go proxy(client, bufConn)
		proxy(conn, client)

		time.Sleep(500 * time.Millisecond)
		client.Close()
	}

	conn.Close()
}

func proxy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Println("Porxy: ", err.Error())
	}
}

func main() {
	l, err := net.Listen("tcp", ":3100")
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}

		go handleRequest(conn)
	}
}
