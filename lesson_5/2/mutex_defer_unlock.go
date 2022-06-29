// 2. Реализуйте функцию для разблокировки мьютекса с помощью defer

package main

import (
	"fmt"
	"sync"
	"time"
)

const count = 1000

func main() {
	var (
		counter int
		mutex   sync.Mutex
		ch      = make(chan struct{}, count)
	)
	for i := 0; i < count; i += 1 {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter += 1
			ch <- struct{}{}
		}()
	}
	time.Sleep(2 * time.Second)
	close(ch)
	i := 0
	for range ch {
		i += 1
	}
	fmt.Println(counter)
	fmt.Println(i)

}
