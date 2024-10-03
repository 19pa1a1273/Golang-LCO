package main

import (
	"fmt"
	"sync"
)

// go env -w CGO_ENABLED=1
// Download the MinGW installer from MinGW SourceForge.

func main() {

	fmt.Println("Race Condition -- LearnCodeOnline.in")

	wg := &sync.WaitGroup{}
	// mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	// You are trying to lock the resource directly that is not good
	// always be locking the resource when you are trying to READ that

	// mut.RLock()
	var score = []int{0}
	// mut.RUnlock()

	//wg.Add(1) // instead of wg.Add(1) --> write once for three - wg.Add(3)
	wg.Add(3)

	// If you just use Mutex (not RWMutex). you don't need to add RLock.
	// Read is a quick operation, Compared to write

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("READ form RWMutex")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	// mut.RLock()
	fmt.Println(score)
	// mut.RUnlock()

}
