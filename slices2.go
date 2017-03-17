package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8}
	b := []int{1, 3, 5, 7}

	//appending multiple values to a slice:
	a = append(a, 10, 12, 14)

	fmt.Printf("a is now %v\n", a)

	//appending a slice to a slice:
	a = append(a, b...)

	fmt.Printf("a is now %v\n", a)

	//appending a slice of a slice to a slice of a slice:
	a = append(a[:2], b[:2]...)

	fmt.Printf("a is now %v\n", a)
}
