// 1. С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из
// которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться, что при
// каждом запуске программы итоговое число равно 1000.
// 2. Написать программу, которая при получении в канал сигнала SIGTERM останавливается не
// позднее, чем за одну секунду (установить таймаут).

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- 1
	}
}

func main() {
	numJobs := 1000
	numWorkers := 50
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		defer close(exit)
		sig := <-exit
		panic(sig)
	}()

	for j := 0; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	var n int
	for a := 1; a <= numJobs; a++ {
		n = n + <-results
	}
	close(results)

	fmt.Println(n)
}
