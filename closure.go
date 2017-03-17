package main

import "fmt"

func main() {
	var a, b int

	add := func() int {
		return a + b
	}

	a, b = 3, 9
	fmt.Printf("%v + %v = %v\n", a, b, add())
	a, b = 4, -7
	fmt.Printf("%v + %v = %v\n", a, b, add())
	/*
	   type of 'add' here is a function that takes in no params and returns an int

	   CLOSURES: with functionss declared within other functions, they close over the local variables within the function
	   so 'add' is able to access a and b because it is a closure *** accessible in the scope closure function
	   **because we assign the function 'add' as a closure (aka define the type)
	*/
	// -----------------------------------------------------------------------------------------------------------------------------

	//note: this is within the main() function
	nextFibb := fibb() //when we assign the fibb() func to the nextFibb var, we are calling the nextFibb function, which is the return val of fibb()
	fmt.Print("1 1 ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", nextFibb()) //so here we print out the value of the return value of nextFibb() which will get the anon func returned by the next function, it will alter the values of f and prevf and return those values eventhough they are declared within the other function and are out of the local scope of this function. The closure lets them be available within this scope
	}
}

func fibb() func() uint {
	var prevf uint = 1
	var f uint = 1       //defined in the scope of function fibb
	return func() uint { //return a function that returns an unsigned integer
		f, prevf = f+prevf, f //take this number and the p revious number, and we walk forward changing f so that it is equal to f+prevf and prev f is equal to f
		return f
	} // this becomes the return value for fibb() and this function will execute closing over each of the sequential fib numbers in the series
}

//can return a closure from a returning function and that closure will close over any variables defined within the function in which that closure is defined
