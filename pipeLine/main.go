// PipeLine - когда мы переливаем данные из одного канала в другой
package main

import (
	"fmt"
	"time"
)

// writer — первая стадия pipeline
// Генерирует числа от 1 до 10 и отправляет их в канал
func writer() <-chan int {
	ch := make(chan int) // Создаем канал для передачи int

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1 // Отправляем числа от 1 до 10
		}
		close(ch) // Закрываем канал по завершении отправки
	}()

	return ch // Возвращаем канал только для чтения
}

// double — вторая стадия pipeline
// Получает числа из inputCh, умножает их на 2, и отправляет дальше
func double(inputCh <-chan int) <-chan int {
	ch := make(chan int) // Создаем канал для передачи результатов

	go func() {
		for v := range inputCh { // Читаем входные данные из канала
			time.Sleep(500 * time.Millisecond) // Имитация "тяжелой" обработки
			ch <- v * 2                        // Умножаем значение на 2 и отправляем дальше
		}
		close(ch) // Закрываем выходной канал после обработки всех значений
	}()

	return ch // Возвращаем выходной канал
}

// reader — последняя стадия pipeline
// Читает из канала и выводит данные в консоль
func reader(ch <-chan int) {
	for v := range ch { // Читаем значения из канала до его закрытия
		fmt.Println(v) // Печатаем результат
	}
}

func main() {
	reader(double(writer()))
}
