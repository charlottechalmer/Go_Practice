package main

import "fmt"

type Rect struct {
	w, h float64
}

//created a new type called "Rect" (SHOULD START WITH A CAPITAL LETTER because anything that starts with a capital letter will be exported to other packages)

func main() {

	//r1 is of type Rect
	var r1 Rect
	r1.w = 5.3
	r1.h = 2.7

	fmt.Printf("Area of r1 = %.2f\n", area(r1))

	//r2 is a pointer to a Rect
	r2 := new(Rect)
	r2.w = 3.5
	r2.h = 7.1

	fmt.Printf("Area of r2 = %.2f\n", area(*r2)) //need to use * b/c it is the indirection operator and will call area to the specific place  in memory that the reference r2 has already allocated within the context of the allocation
}

func area(r Rect) float64 {
	return r.w * r.h
}
