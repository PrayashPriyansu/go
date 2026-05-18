package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	fmt.Println(time.Saturday, time.Now().Weekday()+1)
	t := time.Now()
	// a better alternative to else if chain
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
