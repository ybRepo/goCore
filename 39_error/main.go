package main

import (
	"fmt"
	"log"
	"os"
	"math"
	"errors"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	openfile()
	log.Println(sqrt(-5.0))

}

func openfile() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}


func sqrt(f float64) (float64, error){
	err := errors.New("norgate math: square root of negative number")
	if f<0{
		//log.Println(err)					//should i be logging the err at this time or is it better to wrap the method in main with when calling the function sqrt() as seen in line 21?
		return 0, err

	}
		return math.Sqrt(f), nil
}
