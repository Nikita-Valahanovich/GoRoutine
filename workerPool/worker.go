package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}()

	now := time.Now()
	processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))

}

// операция должна выполняться не более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int) {
	var wg sync.WaitGroup

	// запуск numWorkers горутин
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for val := range in {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

				// для получения результата без буфера
				resultCh := make(chan int)
				doneCh := make(chan struct{})

				// запускаем обработку
				go func(v int) {
					result := processData(v)
					select {
					case resultCh <- result:
					case <-ctx.Done():
						// если контекст отменён, результат не нужен
					}
					close(doneCh) // сигнал, что горутина завершилась
				}(val)

				select {
				case res := <-resultCh:
					out <- res
				case <-ctx.Done():
					fmt.Println("timeout for value", val)
				}

				// чтобы не было утечек
				<-doneCh
				cancel()
			}
		}()
	}

	// отдельная горутина для закрытия out после завершения всех воркеров
	go func() {
		wg.Wait()
		close(out)
	}()
}
