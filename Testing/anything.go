package main

import "fmt"

// DECLARE there is a variable with the identifier "z"
// and that the variable with the identifier "z" is of Type int
var z int

// Declear the variable "y"
// Assign the value 43
// declare & assign = initilization
var y = 43

func main() {

	fmt.Printf("%T\n", y)

	fmt.Println("Testing")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("now exited")
}
