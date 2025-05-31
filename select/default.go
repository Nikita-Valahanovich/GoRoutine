package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch2 <- 1
	}()

	// select - блокирующая операция
	// Чтобы сделать select неблокирубщим, добавьте default
	select {
	case v := <-ch1:
		fmt.Println("v =", v, "ch1")
	case v := <-ch2:
		fmt.Println("v =", v, "ch2")
	default:
		fmt.Println("exited by default") // сработает при условии, что предыдущием кейсы заблокированы
	}
}
