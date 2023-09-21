package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello World")

	b := make([]byte, 6) // [0,0,0,0,0,0]
	for {
		n, err := r.Read(b)                               // returns: the number of bytes filled in, an error value
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b) // n=6, err = nil, b = [72,101,108,108,111,32]
		fmt.Printf("b[:n] = %q\n", b[:])                  // b[:6] --> [72,101,108,108,111,32]
		if err == io.EOF {
			break
		}
	}
}
