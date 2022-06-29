// 1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех

package main

import (
	"sync"
)

const count = 1000

func main() {
	var (
		wg = sync.WaitGroup{}
		mu = sync.Mutex{}
	)
	wg.Add(count)
	for i := 0; i <= count; i += 1 {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
