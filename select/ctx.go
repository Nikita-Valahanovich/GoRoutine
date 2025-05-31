package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	timer := time.NewTimer(1 * time.Millisecond) // таймер 1 миллисекунда

	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel() // cancel нужно деферить, чтобы при выходе из функции cancel сработал

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
	case <-ctx.Done(): // ctx имеет свой конал Done()
		fmt.Println("exited by context")
	}
}
