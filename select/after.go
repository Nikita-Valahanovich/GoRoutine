package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// select - блокирующая операция
	// Чтобы сделать select неблокирубщим, добавьте default
	select {
	case v := <-ch1:
		fmt.Println("v =", v, "ch1")
	case v := <-ch2:
		fmt.Println("v =", v, "ch2")
	case <-time.After(1 * time.Second): // выходим через секунду
		fmt.Println("exited by after") // выводим сообщение
	}
}
