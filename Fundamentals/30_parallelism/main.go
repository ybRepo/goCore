package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup


func main() {
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(3))

	t := time.Now()
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
	wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Bar:", i)
	}
}
