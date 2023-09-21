package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type rot13Reader struct {
	r2 io.Reader
}

type ascii struct {
	a io.Reader
}

func (a ascii) Read(b []byte) (int, error) {
	n, err := a.a.Read(b)
	fmt.Println("---", b[:n])

	return n, err

}

func (r1 rot13Reader) Read(b []byte) (int, error) {

	n, err := r1.r2.Read(b)   // returns: the number of bytes filled in, an error value
	fmt.Println("---", b[:n]) // [76 98 104 32 112 101 110 112 120 114 113 32 103 117 114 32 112 98 113 114 33]

	for i := 0; i < n; i++ {
		switch {
		case b[i]+13 > 'z':
			in := rune(b[i])
			fmt.Printf("%v = %c ", b[i], in)
			b[i] -= 13
			character := rune(b[i])
			fmt.Printf("%v = %c ", b[i], character)
		case b[i] > 'a':
			in := rune(b[i])
			fmt.Printf("%v = %c ", b[i], in)
			b[i] += 13
			character := rune(b[i])
			fmt.Printf("%v = %c ", b[i], character)
		case b[i]+13 > 'Z':
			in := rune(b[i])
			fmt.Printf("%v = %c ", b[i], in)
			b[i] -= 13
			character := rune(b[i])
			fmt.Printf("%v = %c ", b[i], character)
		case b[i] > 'A':
			in := rune(b[i])
			fmt.Printf("%v = %c ", b[i], in)
			b[i] += 13
			character := rune(b[i])
			fmt.Printf("%v = %c ", b[i], character)
		}
	}
	return n, err
}

func main() {
	alpha := strings.NewReader(Alphabet)
	a := ascii{alpha}
	io.Copy(os.Stdout, &a)

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)

}
