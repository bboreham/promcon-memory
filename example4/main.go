package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	mb := 100
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Could not parse %q: %v\n", os.Args[1], err)
			os.Exit(1)
		}
		mb = i
	}

	big := make([]byte, mb*1024*1024)

	for i := range big {
		big[i] = 42
	}

	file, err := os.Create("example4.data")
	if err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(1)
	}
	buf := make([]byte, 4096)
	for {
		file.Write(buf)
		time.Sleep(1 * time.Millisecond)
	}

	runtime.KeepAlive(big)
}
