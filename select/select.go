package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		wg.Done()
		ch2 <- 1
	}()

	wg.Wait()

	// select - блокирующая операция
	// Чтобы сделать select неблокирубщим, добавьте default
	select {
	case v := <-ch1:
		fmt.Println("v =", v, "ch1")
	case v := <-ch2:
		fmt.Println("v =", v, "ch2")
	case <-time.After(1 * time.Second): // выходим через секунду, сработает без wg
		fmt.Println("exited by after") // выводим сообщение
	default:
		fmt.Println("exited by default") // сработает при условии, что предыдущием кейсы заблокированы
	}
}
