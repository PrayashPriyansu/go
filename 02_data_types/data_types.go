package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 4 types of datatypes
// - basic(number,string, bool)
// - aggregate(array, struct)
// - reference(pointers, slices, functions, channel, maps)
// - interfaces

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	name   string     = "Prayash"
	price  float64    = 15.256
)

// bool
// zero value= false
//
// string
// zero value = ""
//
// int
// zero value = 0
// int8 = -128 to 127, 8 bits
// int = 64 bit, unit = 64 bits(based on system architecture)
// generally recommended to go with just normal int and uint
// but use uint8 or others if you know the range of data is definitely less
// byte is an alias for uint8
// rune is an alias for int32, represents the Unicode code point
//
// float32 and float64
// zero value = 0
//
// complex64 and complex128

const Pi = 3.14 // used to create constants

func main() {
	display()

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// there is no implicit conversion
	// always use explicit conversion

	fmt.Println(x, y, z)

	v := 34
	// type infernce
	// it automatically determines the type
	// once inferred can not change
	fmt.Println(v)
}

func display() {
	fmt.Printf("%T - %t\n", ToBe, ToBe)
	fmt.Printf("%T - %d\n", MaxInt, MaxInt)
	// DEFAULT FORMAT
	fmt.Printf("%T - %v\n", z, z)
	// STRING
	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name) // string with double quotes

	// FLOAT
	fmt.Printf("%f\n", price)

	// FLOAT WITH PRECISION
	fmt.Printf("%.2f\n", price)

	// // STRUCT DEBUGGING
	// fmt.Printf("%+v\n", struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "Prayash",
	// 	Age:  21,
	// })

	// // GO-SYNTAX REPRESENTATION
	// fmt.Printf("%#v\n", age)

	// // MEMORY ADDRESS
	// fmt.Printf("%p\n", &age)

	// // PERCENT SYMBOL
	// fmt.Printf("100%% Complete\n")
}
