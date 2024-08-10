package goroutines

import "time"

func someTask(id int, data chan int) {
	for taskId := range data {
		time.Sleep(2 * time.Second)
		println("Task", id, ":", taskId)
	}
}

func Channels() {
	channel := make(chan int)

	// Creating 10.000 workers to execute the task
	for i := 0; i < 10000; i++ {
		go someTask(i, channel)
	}

	// Filling channel with 100.000 numbers to be executed
	for i := 0; i < 100000; i++ {
		channel <- i
	}
}
