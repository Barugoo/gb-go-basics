package main

import (
	"log"
	"os"
	"time"
)

func main() {
	cancel := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel <- 1
	}()

	log.Println("Countdown started. Hit [Enter] to cancel")
	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case moment := <-tick:
			log.Println(moment.Format("15:04:05"), countdown)
		case <-cancel:
			log.Println("Launch canceled!")
			return
		}
	}
	log.Println("The rocket starts!")
}
