package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First         string //unlike prev examples of structs, in order to export the field values, the first letter needs
	Last          string //to be capitalized so that it is visible outside of the package.
	Middle        string `json:"-"` //the field also contains a tag which removes it from the Marchal JSON
	Age           int    `json:"wisdom-score:"`
	isnotexported int
}

func main() {
	p1 := person{"James", "Bond", "James", 20, 0}
	bs, _ := json.Marshal(p1) //Marshal returns the JSON encoding of p1 in []Bytes
	fmt.Println(p1.First, p1.Last, p1.Middle, p1.Age, p1.isnotexported)
	fmt.Println(bs)
	fmt.Println(string(bs)) //to get the string, we need to convert the Marchal encoding of p1 to string

	unmarshall()

}

func unmarshall() {
	var p2 person
	bs := []byte(`{"First":"James","Last":"Bond","wisdom-score:":20}`) //notice wisdom-score: is in the json coming back and when assigning the values to the person fields, the tags allowed Go to know the value 20 represented Age in the person struct
	json.Unmarshal(bs, &p2)                                            //Json values stored in bs are unmarshalled to the addresses of the fields in person struct
	fmt.Println(p2.First, p2.Last, p2.Middle, p2.Age, p2.isnotexported)
}
