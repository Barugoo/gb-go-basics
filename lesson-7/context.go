package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
)

func worker(ctx context.Context, mu *sync.RWMutex, m map[string]string, wg *sync.WaitGroup, num int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			mu.RLock()
			log.Println(m["worker %d doing some work"])
			mu.RUnlock()

			mu.Lock()
			m["worker %d doing some work"] = fmt.Sprintf("worker %d doing some work", num)
			mu.Unlock()
		}
	}
}

const numOfWorkers = 10

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel()
	}()

	m := map[string]string{}
	mu := &sync.RWMutex{}

	wg := &sync.WaitGroup{}
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(ctx, mu, m, wg, i+1)
	}
	wg.Wait()
	log.Println("wg waited")
}
