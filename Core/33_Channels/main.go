package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func main() {
	//	wg.Add(2)

	c := make(chan int)

	go func() {
		//		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("add: ", i)
			c <- i
		}

	}()

	go func() {
		//		defer wg.Done()
		for {
			x := <- c
			fmt.Println("remove: ", x)
		}
	}()

	//	wg.Wait()
	time.Sleep(time.Millisecond)
}
