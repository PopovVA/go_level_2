// 1. Для закрепления навыков отложенного вызова функций, напишите программу, содержащую
// вызов функции, которая будет создавать паническую ситуацию неявно. Затем создайте
// отложенный вызов, который будет обрабатывать эту паническую ситуацию и, в частности,
// печатать предупреждение в консоль. Критерием успешного выполнения задания является то,
// что программа не завершается аварийно ни при каких условиях.
// 2. Дополните функцию из п.1 возвратом собственной ошибки в случае возникновения панической
// ситуации. Собственная ошибка должна хранить время обнаружения панической ситуации.
// Критерием успешного выполнения задания является наличие обработки созданной ошибки в
// функции main и вывод ее состояния в консоль

package homework

import (
	"fmt"
	"math/rand"
	"time"
)

type ErrorWithTime struct {
	text string
	time string
}

func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now().GoString(),
	}
}
func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\ntime: %s", e.text, e.time)
}

func PanicRecover() {
	defer func() {
		if v := recover(); v != nil {
			err := New("Recovered")
			fmt.Println(err)
		}
	}()
	fakePanic()
}

func fakePanic() {
	rand.Seed(time.Now().UnixNano())
	panicValue := 1
	randomValue := rand.Intn(0 + 3)
	if randomValue == panicValue {
		panic("PANIC")
	} else {
		fmt.Println("No panic")
	}

}
