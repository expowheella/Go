package main

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

var s string = "I love McDonalds, their meals and burgers and Coca Cola"

func WordCount(s string) map[string]int {
	result := strings.Fields(s)

	map0 := make(map[string]int)

	var count int = 1

	for i := 0; i < len(result); i++ {
		_, yes := map0[result[i]]
		map0[result[i]] = count
		if yes {
			map0[result[i]] += count
		} else {
			map0[result[i]] = 1
		}
		count = 1
	}

	return map0
}

func main() {
	//wc.Test(WordCount)
	fmt.Print(WordCount(s))
}
