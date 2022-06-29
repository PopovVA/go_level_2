// 1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех

package main

import (
	"fmt"
	"sync"
)

const count = 1000

func main() {
	var (
		wg = sync.WaitGroup{}
	)
	wg.Add(count)
	for i := 0; i <= count; i += 1 {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
