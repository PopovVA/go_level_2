// 4. Выполните задание из блока 'Для самостоятельного изучения' данной методички

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer func() {
// 		if v := recover(); v != nil {
// 			fmt.Println("recovered", v)
// 		}
// 	}()
// 	go func() {
// 		panic("A-A-A!!!")
// 	}()
// 	time.Sleep(time.Second)
// }

// Предложите реализацию примера так, чтобы аварийная
// остановка программы не выполнилась.

package homework

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered", v)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}
