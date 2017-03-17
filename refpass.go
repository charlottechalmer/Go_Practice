package main

import "fmt"

// func main() {
// 	num := 5
// 	num = negate(num)
// 	fmt.Printf("num is now %v\n", num)
// }

// func negate(a int) int {
// 	return a * -1
// } //within this function, we are passing by value aka passing a copy of that value rather than altering and actually changing the original value
//**** IF WE WANT TO PASS BY REFERENCE WE NEED TO USE A POINTER!:

func main() {
	num := 5
	negate(&num) //takes the address of num, passes it in as a pointer to an int, multiples by -1, and then num has actually been changed
	fmt.Printf("num is now %v\n", num)
} //!!! PASSING THE ACTUAL ADDRESS

func negate(a *int) { //use *int because then it is a pointer to that int. dont need the second int because we arent returning anything
	*a *= -1
}
