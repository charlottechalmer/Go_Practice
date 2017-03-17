package main

import "fmt"

func main() {
	//make a sliuce of int with an initial length of 10 and a capacity of 100
	a := make([]int, 10, 100)

	for i := 0; i < 10; i++ {
		a[i] = 2 * i
	}
	fmt.Printf("a is %v\n", a)

	fmt.Printf("adding an element to a...\n")

	a = append(a, 20)

	fmt.Printf("a is now %v\n", a)

	b := a[1:5] //a "slice" of the slice a
	fmt.Printf("b is %v\n", b)

}
