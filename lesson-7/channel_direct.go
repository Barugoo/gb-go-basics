package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Countdown started.")
	tick := make(<-chan time.Time)    // создаем однонаправленный канал
	tick = time.Tick(1 * time.Second) // создаем поток секундных «тиков»
	for countdown := 10; countdown > 0; countdown-- {
		moment := <-tick
		fmt.Println(moment.Format("15:04:05"), countdown)
	}
	fmt.Println("The rocket starts!")
}
