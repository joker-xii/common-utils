package common_utils

import (
	"math/rand"
	"sync"
)

func MergeTwoMap(a, b map[string]int) map[string]int {
	ret := map[string]int{}
	for k, v := range a {
		ret[k] += v
	}
	for k, v := range b {
		ret[k] += v
	}
	return ret
}

func MergeMaps(a ...map[string]int) (ret map[string]int) {
	before := a
	for {

		sz := len(before)
		sz_after := int(sz / 2)

		after := make([]map[string]int, sz_after)

		wg := sync.WaitGroup{}
		wg.Add(sz_after)

		for i, _ := range after {
			go func(k int) {
				after[k] = MergeTwoMap(before[k], before[sz_after+k])
				wg.Done()
			}(i)
		}

		wg.Wait()

		if sz%2 != 0 {
			after[0] = MergeTwoMap(after[0], before[sz-1])
		}

		before = after

		if sz_after == 1 {
			ret = after[0]
			break
		}

	}
	return
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
