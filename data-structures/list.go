/*	Note: Initially when I wrote this, I wanted to create a list as a wrapper around the array, which would dynamically
	resize. However, as a novice, I don't know how to bypass the go compiler that so desperately requires a constant in
	the array constructor (At least without writing a shit ton of switch statements).

	However, I do know of where I can the answer I want: https://golang.org/src/runtime/slice.go
	At my current level, it is complicated, so I'll revisit this when I actually get proficient enough. It will give me
	so practice with the low-level programming stuff.

	Meanwhile, as a novice, I'll bypass this exercise and focus on data structures and algorithms.
*/
package main

import (
	"fmt"
)

type List struct {
	container []int
	len       int
}

// One method of constructors in Go..?
func NewList(nums ...int) *List {
	newListLen := len(nums)
	newSlice := make([]int, newListLen)
	for i, num := range nums {
		newSlice[i] = num
	}
	return &List{newSlice, newListLen}
}

func main() {
	newList := NewList(1, 2, 3, 4, 5)
	fmt.Println(newList)
}
