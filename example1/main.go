package main

import (
	"os"
)

func main() {
	buf := make([]byte, 1024*1024*1024)

	os.Stdout.WriteString("Hit enter to quit:")
	os.Stdin.Read(buf)
}
