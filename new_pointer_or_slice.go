package main

import "fmt"

// slice 是用make去构造的，而不是用new

func main() {
	IntPointer := new([]int)
	IntSlice := make([]int, 0)
	fmt.Printf("%T, %T\n", IntPointer, IntSlice)
}
