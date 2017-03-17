package main

import "fmt"

func main() {
	//initialize an array of test scores
	scores := [8]float64{87.3, 92.0, 100, 78.9, 84, 98, 72.6, 64}
	//print to the array
	fmt.Printf("Scores: %v\n", scores)

	//calculate the average:
	avg := 0.0 //this will be inferred to be a float64

	//add the scores
	// for i := 0; i < 8; i++ {
	// 	avg += scores[i]
	// }
	for i := range scores {
		avg += scores[i]
	}

	//Can also be written like a while loop:
	//i := 0
	//for i < 8 {
	//  DO WORK
	// i++
	//}

	//divide the sum by 8:
	avg = avg / float64(len(scores))

	//print the average
	fmt.Printf("Average score: %.2f\n", avg)
}
