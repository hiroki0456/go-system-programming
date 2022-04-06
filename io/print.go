package main

import (
	"os"
)

func print() {
	// fmt.Println("Hello World!")
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
