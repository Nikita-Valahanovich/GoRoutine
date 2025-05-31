package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel() // cancel нужно деферить, чтобы при выходе из функции cancel сработал

	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 1000; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("v", v)
		case <-ctx.Done():
			fmt.Println("context timeout")
			return
		}
	}
}
