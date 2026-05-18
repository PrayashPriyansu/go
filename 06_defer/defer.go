package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("pryaash")
	fmt.Println("hello")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

// helllo
// pryaash
// world
//
// It gives this way cause there is basically a stack where stuff and pushed and poped
// And stack is LIFO
