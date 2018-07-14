package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
	if stu == nil {
		fmt.Println("stu is nil.")
	}
}

func live() People {
	var stu *Student
	return stu
}

func main() {
	si := live()
	if si == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	si.Show()
}
