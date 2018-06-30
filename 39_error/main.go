package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
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

	x, err := sqrt(25)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(x)
	}

	y, err := negativeVal(-5)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(y)
	}

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

func sqrt(f float64) (float64, error) {
	ErrCustom := errors.New("norgate math: square root of negative number")
	if f < 0 {

		return 0, ErrCustom

	}
	return math.Sqrt(f), nil
}

func negativeVal(y int) (int, error) {
	ErrNegativeVal := fmt.Errorf("This err includes more details as it includes the negative number: %v", y)
	if y < 0 {
		return 0, ErrNegativeVal
	}
	return y, nil
}
