// 3. Смоделировать ситуацию 'гонки', и проверить программу на наличии 'гонки'

package main

import (
	"fmt"
)

const count = 1000

func main() {
	var (
		counter int
	)

	for i := 0; i < count; i += 1 {
		go func() {
			counter += 1
		}()
	}
	for i := 0; i < count; i += 1 {
		go func() {
			counter += 1
		}()
	}
	fmt.Println(counter)
}
