package main

import (
"fmt"
)

type People2 interface {
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
	var peo People2 = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}