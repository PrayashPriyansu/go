package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)

	// can also directly create a slice
	var s []int          // nil slice
	s2 := []int{1, 2, 3} // non-nil slice
	s3 := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}

	fmt.Println(s2, s3)

	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	for i := range 5 {
		s = append(s, i)
		fmt.Println("len:", len(s), "cap:", cap(s))
	}
}

/*
A slice is a dynamic, flexible view over an underlying array.

INTERNALS:
A slice contains:
- pointer to underlying array
- length
- capacity

Slices DO NOT store data themselves.
They reference an underlying array.

-----------------------------------
DECLARATION
-----------------------------------

nums := []int{1,2,3}

[]int means:
"slice of ints"

-----------------------------------
ARRAY VS SLICE
-----------------------------------

Array:
	var arr [5]int

- fixed size
- size is part of type

Slice:
	var nums []int

- dynamic/flexible
- used far more commonly

-----------------------------------
SLICING
-----------------------------------

arr := [5]int{10,20,30,40,50}

s := arr[1:4]

Result:
[20 30 40]

Rules:
- start inclusive
- end exclusive

-----------------------------------
LENGTH
-----------------------------------

len(s)

Number of accessible elements. [20, 30, 40] - 3

-----------------------------------
CAPACITY
-----------------------------------

cap(s)

Number of elements available from slice start until end of underlying array. [20, 30, 40, 50] - 4

-----------------------------------
APPEND
-----------------------------------

nums = append(nums, 4)

IMPORTANT:
append may allocate a NEW array.

Always reassign append result.

-----------------------------------
SHARED MEMORY
-----------------------------------

Slices share underlying arrays.

Changing slice data may affect:
- original array
- other slices

Example:

a := arr[0:2]
b := arr[1:4]

Both reference same array memory.

-----------------------------------
MAKE
-----------------------------------

make([]int, len, cap)

Example:

nums := make([]int, 3, 10)

len = 3
cap = 10

-----------------------------------
NIL SLICE
-----------------------------------

var nums []int

Valid nil slice:
- len = 0
- cap = 0
- append still works

-----------------------------------
COPY
-----------------------------------

dst := make([]int, len(src))
copy(dst, src)

Used to avoid shared-memory side effects.

-----------------------------------
RANGE LOOP
-----------------------------------

for i, v := range nums {
	fmt.Println(i, v)
}

Ignore index:
for _, v := range nums

-----------------------------------
IMPORTANT MENTAL MODEL
-----------------------------------

Slice = lightweight descriptor over an underlying array.

Passing slices to functions does NOT copy entire data.

Only:
- pointer
- len
- cap

are copied.

*/
