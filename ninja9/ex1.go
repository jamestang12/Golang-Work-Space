package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("being cpu", runtime.NumCPU())
	fmt.Println("bein gs", runtime.NumGoroutine())
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("hello")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

}
