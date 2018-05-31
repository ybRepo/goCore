package visibility

import "fmt"

//PrintVar is capitalized and therefore visible to the rest of the packages
func PrintVar() {
	fmt.Println(yourtitle)
}
