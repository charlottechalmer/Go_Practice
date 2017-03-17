package main

import "fmt"

func main() {
	a, b, c, d, e := 1, 2, 3, 4, 5

	//add through e:
	sum := add(a, b, c, d, e)
	fmt.Printf("The sum is %v\n", sum)
}

func add(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
