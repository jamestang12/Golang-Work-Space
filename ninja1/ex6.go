package main

import (
	"fmt"
)
type hotdog int

var x hotdog
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x) + 2
	fmt.Println(y)

}
