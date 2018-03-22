package main

import (
	"sync"
	"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	m := new(sync.Mutex)
	var ages = UserAges{make(map[string]int,0), *m}
	for i:= 0; i < 1000; i++{
		go ages.Add("aaa"+string(i),i)
		//fmt.Println(i)
	}
	for i:= 0; i < 10; i++{
		fmt.Println(ages.Get("aaa"+string(i)))
		//fmt.Println(i)
	}
	fmt.Println("end")
}