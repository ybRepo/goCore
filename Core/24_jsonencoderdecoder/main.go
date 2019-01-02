package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	encodejson()
	decodejson()

}

func encodejson() {
	p1 := person{"James", "Bond", 20}
	json.NewEncoder(os.Stdout).Encode(p1)
}

func decodejson() {
	var p2 person
	rdr := strings.NewReader(`{"First":"James","Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p2)

	fmt.Println(p2.First, p2.Last, p2.Age)
}
