package main

import "fmt"

// go only has one loop, i.e, for loop
// if you wanna use while loop
// for i<=10
func main() {

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}
