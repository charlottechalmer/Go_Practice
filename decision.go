package main

import "fmt"

func main() {
	a, b := 1, 2

	if a == b {
		fmt.Printf("%v and %v are equal\n", a, b)
	} else if a > b {
		fmt.Printf("%v is greater than %v\n", a, b)
	} else if a < b {
		fmt.Printf("%v is less than %v\n", a, b)
	} else {
		fmt.Printf("%v and %v are not equal\n", a, b)
	}
}
