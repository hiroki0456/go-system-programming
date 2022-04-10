package main

import (
	"fmt"
	"os"
	"time"
)

func fprintf() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
