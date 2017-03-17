package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

/*
	Go does not have constructors or initalizers, but we can return a pointer to a structure to mimic this behavior
*/

func makePoint(a, b float64) *point {
	p := new(point)
	p.x, p.y = a, b
	return p
}

//a method of point:
func (p point) distanceTo(a, b float64) float64 { //before the name of the method, tell the function what it is a member of where p is a point (refer to p inside the function and point is its type))
	dist := math.Sqrt(math.Pow(a-p.x, 2) + math.Pow(b-p.y, 2)) //where Pow is the power
	return dist
}

func main() {
	x, y := 0.0, 0.0

	//get coordinates for a new point
	fmt.Print("Enter x and y coordinates of point, in the form x, y: ")
	fmt.Scanf("%v, %v", &x, &y)

	//myPoint is a point value
	//myPoint := point{x, y}
	//myPoint is a pointer to a point:
	myPoint := &point{x, y}
	//call a function that returns a pointer to a point:
	//myPoint := makePoint(x, y)

	//!!WANT TO BE ABLE TO ASSIGN POINTERS TO DATASTRUCTURES AND THEN USE POINTERS RATHER THAN USE THE ACTUAL VALUES. can use values if only going to have a couple of this data point

	//print myPoint's x, y coordinates:
	fmt.Printf("Point is at (%v, %v)\n", myPoint.x, myPoint.y)

	//print the distance of myPoint from (0, 0)
	fmt.Printf("Distance of myPoint from (0, 0) is %.2v\n", myPoint.distanceTo(0, 0))
}
