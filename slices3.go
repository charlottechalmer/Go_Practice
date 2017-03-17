package main

import "fmt"

func main() {

	//a string is an array of bytes
	testString := "This is a test."

	fmt.Printf("The first word of testString is \"%s\"\n", firstWord(testString))

}

func firstWord(str string) []byte {
	//word is an empty slice of bytes
	word := []byte{}

	//get everything in the argument upto but not including the first space
	for ch := range str {
		if str[ch] == ' ' {
			//break out of the for loop:
			break
		} else {
			//append this byte to word:
			word = append(word, str[ch])
		}
	}

	//return the first word
	return word
}
