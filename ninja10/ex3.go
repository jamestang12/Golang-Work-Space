package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		// close so no race, so to wait for this than no deadlock will happend
		close(c)
	}()

	return c
}
