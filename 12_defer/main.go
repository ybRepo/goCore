package main

import "fmt"

func hello(){
	fmt.Print("Hello ")
}

func world(){
	fmt.Println("World")
}

func main(){
	defer world()		//defer is a way to write code that executes right before the function closes
	hello()				//this is a great way to for example open a file, defer the closing of the file, and then write the code needed
						//to work with the file. This improves your ability to troubleshoot the issue.
}