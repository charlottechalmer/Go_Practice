package main

import (
	"fmt"
)

func main() {

	/*
		a map is a table of key/value pairs. maps are also known as hash tables, dictionaries, or associative arrays in other languages.

		the keys of a map must be "hashable", meaning that a unique value can be assigned to each key and a unique value can be assigned to each key and that equality can be evaluated for every key.add

		ints, floats, strings, etc. are all hashable types, but arrays slices, etc. are not.add

		the values of a map need not be hashable.
	*/

	digits := map[string]int{ //declaring and initializing
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	digits["five"] = 5 //adding a new key/value pair

	if digits["four"] != 0 {
		fmt.Printf("that digit is %v\n", digits["four"])
	} else {
		fmt.Printf("that digit is not in the map")
	}

	//this is called the comma ok idiom. It tests for the presence of a pair in the map
	key := "zero"

	var digit int
	var ok bool

	if digit, ok = digits[key]; ok { //NOTICE THE SYNTAX HERE--> semicolon splits line! assigning the var ok before use so can use it
		fmt.Printf("%v is %v\n", key, digit)
	} else {
		fmt.Printf("%v is not a key in the map. \n", key)
	}

	//we can also throw away the value is we're only interested in knowing if the key exists. this is done to avoid the problem of initalizing but not using a variable
	if _, ok = digits["nine"]; ok {
		fmt.Println("nine is in the map")
	} else {
		fmt.Println("nine is not in the map")
	}

	//deleting a pair from a map:
	delete(digits, "zero")

	//we can print the entire map:
	fmt.Printf("%v\n", digits)

}
