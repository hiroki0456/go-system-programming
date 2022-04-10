package main

import (
	"io"
	"net"
	"os"
)

func netDial() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: exmaple\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
