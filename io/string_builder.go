package main

import (
	"fmt"
	"strings"
)

func stringBuilder() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}
