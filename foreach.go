package main

import "fmt"

type student struct {
	age  int
	name string
}

func main() {
	stus := []student{
		{20, "small"},
		{30, "middle"},
		{40, "large"},
	}
	ms := make(map[string]*student)
	for index, stu := range stus {
		ms[stu.name] = &stu
		fmt.Println(index)
	}

	for key, value := range ms {
		fmt.Printf("key = %s, age = %d, name = %s\n", key, value.age, value.name)
	}
}
