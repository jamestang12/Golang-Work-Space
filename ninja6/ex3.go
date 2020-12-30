package main

import "fmt"

func main() {

	defer foo()
	fmt.Println("Hello sir")

}

func foo() {
	fmt.Println("foo ran")
}
