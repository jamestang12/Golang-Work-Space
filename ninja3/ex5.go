package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("When %v is divied by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}

}
