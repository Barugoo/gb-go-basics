package main

import (
	"fmt"
)

func main() {
	ourSlice := []int{98, 93, 77, 82, 83}

	total := 0
	for _, v := range ourSlice {
		total += v
	}
	fmt.Println(total / int(len(ourSlice)))
}
