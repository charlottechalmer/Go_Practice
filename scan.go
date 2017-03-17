package main

import "fmt"

//IMPLEMENTING SCAN()
// func main() {
// 	word1, word2 := "", ""
// 	fmt.Print("Enter two words: ")
// 	count, err := fmt.Scan(&word1, &word2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("You entered %v. There are %v words.\n", word1+" "+word2, count) //concat like in JS
// }

/*
	with scan, have to use pointers (addresses of destinations)
	the scan family of functions will return a count and an error where count is the number of items that it scanned
	the scan is going to treat new lines as spaces and is going to expect you to put a newline or a space between your choices
*/

//IMPLEMENTING SCANLN()
func main() {
	word1, word2 := "", ""
	fmt.Print("Enter two words: ")
	count, err := fmt.Scanln(&word1, &word2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("You entered %v. There are %v words.\n", word1+" "+word2, count) //concat like in JS
}

/*
	with scanln, new lines are not treated as spaces.
	the newline needs to be after the last argument
*/

/*
	USE WHEN LOOKING FOR DISCRETE VALUES
	!!IF THE USER GIVES MORE THAN THE INDICATED ARGS, THE EXTRA ARGS WILL BE PASSED ONTO THE COMMAND LINE!
*/
