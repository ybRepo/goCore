package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}

func foo() {
	defer wg.Done()
	for i := 0; i < 40; i++ {
		mutex.Lock()
		counter++
		fmt.Println("Counter:", counter)
		mutex.Unlock()
	}
}

func bar() {
	defer wg.Done()
	for i := 0; i < 40; i++ {
		mutex.Lock()
		counter++
		fmt.Println("Counter:", counter)
		mutex.Unlock()
	}
}
