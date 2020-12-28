package main

import "fmt"

func main() {

	bd := 1985
	for bd <= 2017 {
		if bd > 2017 {
			break
		}

		fmt.Println(bd)
		bd++
	}

}
