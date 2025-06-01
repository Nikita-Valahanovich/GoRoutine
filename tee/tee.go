package main

import (
	"fmt"
	"sync"
)

func tee(in <-chan int) (_, _ <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		defer close(out1)
		defer close(out2)

		for v := range in {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case out1 <- v:
					out1 = nil
				case out2 <- v:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}

// генератор
func generate() chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	ch1, ch2 := tee(generate())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch1 {
			fmt.Println("channel 1 = ", v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Println("channel 2 = ", v)
		}
	}()

	wg.Wait()
}
