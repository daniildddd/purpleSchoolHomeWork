package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- rand.Intn(101)
		}
		close(ch)
		wg.Done()
	}()
	go func() {
		for i := range ch {
			fmt.Printf(" %d", i*i)
		}
		wg.Done()
	}()
	wg.Wait()
}
