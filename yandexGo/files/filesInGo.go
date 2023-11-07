package files

import (
	"fmt"
	"os"
)

func Files() {
	// data, err := os.ReadFile("literature.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)

	// Seek
	// opening a file
	file, _ := os.Open("literature.txt")

	// creating a buffer
	buffer := make([]byte, 100)

	// offset value
	var offset int64 = 1024

	// make offset from the beginning of a file
	file.Seek(offset, 0)

	// read the buffer starting from the offset number
	text_length, err := file.Read(buffer)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(buffer))
	fmt.Println(text_length)

}
