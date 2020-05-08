package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	// форматирование
	fmt.Println(t.Format("02-01-2006"))
	fmt.Println(t.Format("02-01-2006 15:04:05"))

	start := time.Now()
	time.Sleep(100 * time.Millisecond) // задержка 100 ms
	end := time.Now()
	fmt.Println(end.Sub(start))
}
