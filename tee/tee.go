// Паттерн Tee
package main

import (
	"fmt"
	"sync"
)

// tee - паттерн дублирования данных из входного канала в 2 других
func tee(in <-chan int) (_, _ <-chan int) {
	out1 := make(chan int) // создаем выходной канал №1
	out2 := make(chan int) // создаем выходной канал №2

	go func() {
		defer close(out1) // закрываем канал после завершения работы функции
		defer close(out2) // закрываем канал после завершения работы функции

		for v := range in { // ренжом получаем данные из входного канала
			var out1, out2 = out1, out2 // делаем затенение
			for i := 0; i < 2; i++ {
				select {
				case out1 <- v: // записываем данные в канал 1
					out1 = nil // блокируем этот канал после записи
				case out2 <- v: // записываем данные в канал 2
					out2 = nil // блокируем этот канал после записи
				}
			}
		}
	}()

	return out1, out2 // возвращаем данные из каналов
}

// generate - паттерн генератора
func generate() chan int {
	ch := make(chan int) // создаем канал, в которой записываем данные

	go func() {
		defer close(ch) // закрываем канал, после завершения goroutine

		for i := 1; i <= 5; i++ {
			ch <- i // записываем значение в канал (от 1 до 5)
		}
	}()

	return ch // возвращаем данные из канала
}

func main() {
	ch1, ch2 := tee(generate())
	wg := &sync.WaitGroup{} // примитив синхронизации

	wg.Add(1)   // создаем счетчик
	go func() { // создаем точку fork, запускаем goroutine
		defer wg.Done() // снижаем счетчик на 1, после завершения работы функции
		for v := range ch1 {
			fmt.Println("channel 1 = ", v)
		}
	}()

	wg.Add(1)   // создаем счетчик
	go func() { // создаем точку fork, запускаем goroutine
		defer wg.Done() // снижаем счетчик на 1, после завершения работы функции
		for v := range ch2 {
			fmt.Println("channel 2 = ", v)
		}
	}()

	wg.Wait() // создаем точку Join, дожидаемся завершения работы горутин
}
