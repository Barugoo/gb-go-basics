package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[0]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	strPtr := flag.String("str", "hello", "a string")
	numPtr := flag.Int("num", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var strVar string
	flag.StringVar(&strVar, "strVar", "world", "a string var")

	flag.Parse()

	fmt.Println("str:", *strPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("strVar:", strVar)
	fmt.Println("tail:", flag.Args())
}
