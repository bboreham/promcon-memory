package main

import (
	"os"
)

func main() {
	buf := make([]byte, 100*1024*1024)

	for i := range buf {
		buf[i] = 42
	}

	os.Stdout.WriteString("Hit enter to quit:")
	os.Stdin.Read(buf)
}
