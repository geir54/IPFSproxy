package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestRequest(t *testing.T) {
	b := []byte("GET / HTTP/1.1\r\n" +
		"Host: arstechnica.com\r\n" +
		"User-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:45.0) Gecko/20100101 Firefox/45.0\r\n" +
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\n" +
		"Accept-Language: en-US,en;q=0.5\r\n" +
		"Accept-Encoding: gzip, deflate\r\n" +
		"Cookie: adfbadbf;:adsfasdf\r\n" +
		"DNT: 1\r\n" +
		"Connection: keep-alive\r\n" +
		"Pragma: no-cache\r\n" +
		"Cache-Control: no-cache\r\n\r\n" +
		"bla:bla")
	reader := bytes.NewReader(b)
	bufConn := bufio.NewReader(reader)

	req := newRequest()
	req.Decode(bufConn)

	if req.Path != "/" {
		t.Fatalf("Path does not match")
	}

	if req.Proto != "HTTP/1.1" {
		t.Fatalf("Proto does not match")
	}

	_, ok := req.Header["bla"]
	if ok {
		t.Fatalf("This should not have been read")
	}

	if req.Header["Host"] != "arstechnica.com" {
		t.Fatalf("Host does not match")
	}
}
