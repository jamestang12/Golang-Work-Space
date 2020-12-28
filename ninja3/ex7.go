package main

import "fmt"

func main() {

	x := "asd"

	if x == "moneypenny" {
		fmt.Println(x)
	} else if x == "testing" {
		fmt.Println("Testing")
	} else {
		fmt.Println("neither")
	}

}
