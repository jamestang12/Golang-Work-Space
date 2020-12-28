package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	x = append(x, 12)
	fmt.Println(x)
	x = append(x, 13, 14, 15)
	fmt.Println(x)

	y := []int{56, 57, 58}
	x = append(x, y...)
	fmt.Println(x)
}
