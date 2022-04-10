package main

import (
	"bytes"
	"fmt"
)

func buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
