package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Printf("typeof peo %T\n", peo)
	fmt.Println(peo.Speak(think))
}
