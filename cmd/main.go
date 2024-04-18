package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

func main() {
	var buff = make([]byte, 4096)
	var data = make([]byte, 0)
	var ticker = time.NewTicker(time.Second)

	for range ticker.C {
		if _, err := rand.Read(buff); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)

		} else {
			data = append(data, buff...)
			fmt.Fprintf(os.Stdout, "Size: %d bytes\n", len(data))
		}
	}
}
