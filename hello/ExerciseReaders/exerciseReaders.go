// it would be better to say: "rewrite all values in []byte into 'A's"

package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{} // empty structure

func (m MyReader) Read(b []byte) (int, error) {

	// b = [65][65][65][65][65][65] with length = 1024 
	// ASCII Code of 'A' = 65

	for x := range b {
		// we just replace each value by 'A'
		b[x] = 'A'							 	// b[0] = 'A', b[1] = 'A' etc...
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

/* Роль io.Reader.Read заключается в записи в заданную область памяти данных, считанных из ее источника.

By convention an io.Reader indicates its end by returning an io.EOF error. If the reader does not return an error, it behaves as an infinite source of data to its consumer which can never detect an exit condition.
*/
