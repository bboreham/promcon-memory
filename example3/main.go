package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("example3.data")
	if err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(1)
	}
	buf := make([]byte, 4096)
	for {
		file.Write(buf)
		time.Sleep(1 * time.Millisecond)
	}
}
