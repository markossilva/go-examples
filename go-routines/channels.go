package goroutines

import "time"

func someTask(id int, data chan int) {
	for taskId := range data {
		time.Sleep(2 * time.Second)
		println("Task", id, ":", taskId)
	}
}

func Channels(numWorkers, numExecutions int) {
	channel := make(chan int)

	for i := 0; i < numWorkers; i++ {
		go someTask(i, channel)
	}

	for i := 0; i < numExecutions; i++ {
		channel <- i
	}
}
