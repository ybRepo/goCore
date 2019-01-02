package main

import (
	"fmt"
)

//var wg sync.WaitGroup

func main() {
	//	wg.Add(2)

	c := make(chan int)
	done := make(chan bool)

	go func() {
		//		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("add: ", i)
			c <- i
		}
		done <-true
	}()

	go func() {
		//		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("add: ", i)
			c <- i
		}
		done <-true
	}()

	go func() {
		<- done
		<- done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}


	semaphore()
}


func semaphore(){
	n :=10
	d := make(chan int)
	done := make(chan bool)

	for i:=0; i<n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				d <- i
			}
			done <- true
		}()
	}

		go func (){
			for i:=0; i < n; i++{
				<- done

			}
			close(d)
		}()

	for n := range d {
		fmt.Println(n)
	}
}


