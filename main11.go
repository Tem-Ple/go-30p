package main

import (
	"fmt"
)

type People3 interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People3 {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}