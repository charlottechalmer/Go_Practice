package main

import "fmt"

func main() {
	a, b := 3, 5

	fmt.Printf("Before swap: a = %v, b = %v\n", a, b)

	//swap a and b:
	a, b = swap(a, b)

	fmt.Printf("After swap: a = %v, b = %v\n", a, b)
}

func swap(a, b int) (x int, y int) {
	x, y = b, a
	return
}
