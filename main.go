package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// sample 1
	//
	// nilness potential
	var points []int
	fmt.Println("show point for index '5'", points[5])

	// sample 4
	//
	// missing error check
	unecessaryNewLine()

}

// sample 2
//
// deadcodes
func sampleDeadcodes() {
	fmt.Print("this function is unused / deadcodes")
}

func unecessaryNewLine() error {

	// sample 3 ^ uncessary new line
	var username string
	fmt.Println("username: ", username)

	return nil
}
