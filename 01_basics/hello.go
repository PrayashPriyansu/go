package main

import (
	"fmt"
	"math/rand"
)

var a, b, c bool // this is how variables are defined outside the function

func main() {
	a, b := swap(14, 6) // this is how variables are defined inside the function

	fmt.Println(a, b)
	fmt.Println(add(5, 7))
	fmt.Println(rand.Intn(10))
	fmt.Println("Hello, 世界")
}

func add(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func split(sum int) (x, y int) { // dont use normally

	x = sum * 4 / 9
	y = sum - x

	return
}
