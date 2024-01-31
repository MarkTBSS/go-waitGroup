package main

import (
	"fmt"
	"sync"
)

func main() {
	sliceA := []int{1, 2, 3, 4, 5}

	var waitGroupVariable sync.WaitGroup
	for i := range sliceA {
		waitGroupVariable.Add(1)
		go func(i int) {
			defer waitGroupVariable.Done()
			fmt.Printf("%v ", sliceA[i])
		}(i)
	}
	waitGroupVariable.Wait()
}
