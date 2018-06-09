package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	t := time.Now()

	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	fmt.Println(time.Since(t))

}

func foo() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Bar:", i)
	}
}
