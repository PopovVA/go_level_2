// 1. Написать программу, которая использует мьютекс для безопасного доступа к данным
// из нескольких потоков. Выполните трассировку программы

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

const count = 1000

func main() {
	var (
		counter int
		lock    sync.Mutex
		wg      sync.WaitGroup
	)
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg.Add(count * 2)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			lock.Lock()

			defer lock.Unlock()
			counter += 1
		}()
	}
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			lock.Lock()

			defer lock.Unlock()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
