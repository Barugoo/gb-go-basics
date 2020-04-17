package main

import "fmt"

const pi = 3.141

const (
	hello = "Привет"
	e     = 2.718
)

const (
	// нетипизированная константа
	year = 2017
	// типизированная константа
	yearTyped int = 2017
)

func main() {
	var month int32 = 13
	fmt.Println(month + year)

	// month + yearTyped (mismatched types int32 and int)
	fmt.Println(month + int32(yearTyped))
}
