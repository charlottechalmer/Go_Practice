package main

import "fmt"

func main() {

	// grades := [7]int{93, 72, 68, 54, 100, 43, 86}
	// strScore := ""
	// for score := range grades {
	// 	switch {
	// 	case grades[score] > 90:
	// 		strScore = "A"
	// 	case grades[score] > 80:
	// 		strScore = "B"
	// 	case grades[score] > 70:
	// 		strScore = "C"
	// 	case grades[score] > 60:
	// 		strScore = "D"
	// 	default:
	// 		strScore = "F"
	// 	}
	// 	fmt.Printf("%v, letter grade %v\n", grades[score], strScore)
	// }

	//SWITCH AND LOOP W/ VARS:
	// numSize := ""
	//
	// //loop through from 1 to 10
	// for i := 1; i <= 10; i++ {
	// 	//assign a numSize to i
	// 	if i < 4 {
	// 		numSize = "small"
	// 	} else if i < 8 {
	// 		numSize = "medium"
	// 	} else {
	// 		numSize = "large"
	// 	}
	// 	//switch to numSize
	// 	switch numSize {
	// 	case "small":
	// 		fmt.Printf("%v is a small number.\n", i)
	// 	case "medium":
	// 		fmt.Printf("%v is a medium number.\n", i)
	// 	default:
	// 		fmt.Printf("%v is a large number.\n", i)
	// 	}
	//
	// }

	//SWITCH WITH COMBINED CASES:
	//loop from 1 to 10:
	for i := 1; i <= 10; i++ {
		//assign a numSize to i
		switch i {
		case 1, 2, 3:
			fmt.Printf("%v is a small number.\n", i)
		case 4, 5, 6, 7:
			fmt.Printf("%v is not small but not a large number.\n", i)
		default:
			fmt.Printf("%v is a large number.\n", i)
		}
	}

}
