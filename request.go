package main

import (
	"bufio"
	"log"
	"strings"
)

type Request struct {
	Method string
	Path   string
	Proto  string
	Header map[string]string
}

func newRequest() Request {
	return Request{
		Header: make(map[string]string),
	}
}

func (r *Request) Decode(reader *bufio.Reader) error {
	first, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Decode: " + err.Error())
		return err
	}

	sfirst := strings.Split(first, " ")
	r.Method = sfirst[0]
	r.Path = sfirst[1]
	r.Proto = sfirst[2][:len(sfirst[2])-2]

	// Read the header
	var b byte
	var header string
	for err == nil {
		b, err = reader.ReadByte()

		header = header + string(b)
		if len(header) > 4 {
			if header[len(header)-4:] == "\r\n\r\n" {
				header = header[:len(header)-4]
				break
			}
		}
	}

	sheader := strings.Split(header, "\r\n")

	for _, element := range sheader {
		r.Header[element[:strings.Index(element, ":")]] = element[strings.Index(element, ":")+2:]
	}

	return nil
}

func (r *Request) toBytes() []byte {
	b := []byte(r.Method + " " + r.Path + " " + r.Proto + "\r\n")
	for key, val := range r.Header {
		b = append(b, []byte(key+": "+val+"\r\n")...)
	}
	b = append(b, []byte("\r\n")...)
	return b
}
