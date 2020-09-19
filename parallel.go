package common_utils

import (
	"sync"
)

type Task func()
type Task1 func(int)

func RunTasks(tasks ...Task) {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for _, v := range tasks {
		go func(task Task) {
			task()
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func RunTaskFor(maxIterations int, task Task1) {
	wg := sync.WaitGroup{}
	wg.Add(maxIterations)
	for i := 0; i < maxIterations; i++ {
		go func(n int) {
			task(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func RunTasksWithWorker(workerNum, maxIterations int, task Task1) {
	guard := make(chan struct{}, workerNum)
	wg := sync.WaitGroup{}
	wg.Add(maxIterations)
	for i := 0; i < maxIterations; i++ {
		guard <- struct{}{} // would block if guard channel is already filled
		go func(n int) {
			task(n)
			<-guard
			wg.Done()
		}(i)
	}
	wg.Wait()
}