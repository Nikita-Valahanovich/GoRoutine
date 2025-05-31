package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	timer := time.NewTimer(1 * time.Millisecond) // таймер 1 миллисекунда

	// select - блокирующая операция
	// Чтобы сделать select неблокирубщим, добавьте default
	select {
	case v := <-ch1:
		fmt.Println("v =", v, "ch1")
	case v := <-ch2:
		fmt.Println("v =", v, "ch2")
	case <-time.After(1 * time.Second):
		fmt.Println("exited by after")
	case <-timer.C: // таймер имеет свой канал C
		fmt.Println("exited by timer") // выведет сообщение

	}
}
