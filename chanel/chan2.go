package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			ch <- i * 2
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println("v =", v, "worker1")
		}
	}()

	for v := range ch {
		fmt.Println("v =", v, "worker2")
	}
}
