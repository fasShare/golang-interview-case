package main

import "fmt"

// 往另外一个slice里面append另外一个同类型的slice时需要带上...

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	// error s1 = append(s1, s2)
	fmt.Println(s1)
}
