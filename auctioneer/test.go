package main

import (
	"fmt"
	"sync"
)

var maxValue int

func f(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	maxValue++
	m.Unlock()
	wg.Done()
}

func main2() {
	var wg sync.WaitGroup
	var m sync.Mutex

	count := 10
	maxValue = 0
	wg.Add(count)
	for i := 0; i < count; i++ {
		go f(&wg, &m)
	}
	wg.Wait()
	fmt.Println("Max Value:", maxValue)
}
