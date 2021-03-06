package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// печать
	for {
		fmt.Println(<-squares)
	}
}
