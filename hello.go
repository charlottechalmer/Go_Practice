package main

import "fmt"

func main() {
	a, b := 3, 5
	fmt.Printf("hello, world! \n")
	fmt.Printf("%v + %v = %v\n", a, b, add(a, b))
}

func add(n1, n2 int) int {
	return n1 + n2
}
