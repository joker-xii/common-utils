package common_utils

import (
	"math/rand"
	"sync"
)

type Task func()

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

const RAND_MAX = 10000

func RandomSwitch(p int) bool {
	if p == 0 {
		return false
	} else if p == RAND_MAX {
		return true
	}

	return rand.Int()%RAND_MAX < p
}
