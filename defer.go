package main

import "fmt"

func main() {
	defer f1() //here, f1 is deferred until main is finished running
	f2()
}

func f1() {
	fmt.Println("Hello from f1")
}

// func f2() {
// 	fmt.Println("Hello from f2")
// 	for i := 0; i <= 3; i++ {
// 		fmt.Println(i)
// 	}
// }

//so when this is run, f2() is executed and then when main() is finished running, then f1() is run because it is deferred
//this is especially helpful with cleanup
//BE AWARE: when defer executes, every call is pushed onto the stack.

func f2() {
	fmt.Println("Hello from f2")
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

// when i=0, i is pushed onto the stack as 0, then i becomes 1 and 1 is pushed onto the stack, then i becomes 2 and 2 is pushed onto the stack, then i becomes 3 and 3 is pushed onto the stack.
// BUT BECAUSE STACKS ARE FILO, they are then executed in *REVERSE ORDER* --> "Hello from f2", 3, 2, 1, 0, "Hello from f1"
