package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic("error:newline not found")
	}

	fmt.Println(str)
}

/*
	importing bufio => buffered input and output
	os =>
	looking at the reader, it is a new reader that is going to be reading from the file "os.Stdin"
		** NewReader and Stdin are capitalized because the app is going to have to import or export the values!
	bufio wraps and io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O
		**within os, there is a var called Stdin which is a NewFile with a uint pointer that points to syscall (which is a package) that has a Stdin var defined. so this var itself is a wrapper around syscall.Stdin. So this itself points to "/dev/stdin"


		So here, we are getting back from NewReader, which passes in os.Stdin, we are going to get back a pointer to a reader or a bufio.NewReader, that is going to point to a Stdin

		then print a new string and wait for input,
		then call read string on the reader which expects to find a character on which to terminate its reading (here it is `\n`)
		then returns an error only if it gets to the end of the input only if it doesnt find the character passed in (`\n`)

		the panic function will panic the running process and terminate
			takes a string with the error
			if we aren't interested in the error, can pass _ into line 13 (i.e. str, _ := reader.ReadString(`\n`))
*/
