/*
	a pointer is a variable that hold the address in memory of a value

	each byte in memory is assigned an address

	a pointer holds the address information for a memory location

	two operators are used:
		* : the indirection or "contents of" operator
		& : the "address of" operator

	We can make a pointer to a variable by taking its address using the & operator

		a := 5 //a is an int
		b := &a //b is a pointer to an int (a)
		//a == *b (both are 5)

	So, who cares?
		-when we use new to allocate space for a variable, we get a reference to that variable (a pointer)

		-we can pass a pointer as a parameter to a function in order to change that parameter's value outside the function

		-all languages deal with pointers; some languages hide this face very well
*/

package main

import "fmt"

func main() {
	//a is a pointer to an int
	a := new(int)

	//assign 9 as the "contents of" a
	*a = 9

	//print the address and contents of a
	fmt.Printf("a is %v, *a is %v\n", a, *a)

	//b points to the same location as a
	b := a

	//print the address and contents of b
	fmt.Printf("b is %v, *b is %v\n", b, *b)

	//increment the contents of a
	*a++

	//because b references the same memory as a, changing a also changes b:
	fmt.Printf("b is now %v, *b is %v\n", b, *b)
}
