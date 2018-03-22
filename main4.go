package main

import "fmt"

type People struct{a int}

func (p *People) PlusOne1() {
	p.a++
}
func (p People) PlusOne2() {
	p.a++
}

func main() {
	t := People{1}
	fmt.Println(t.a)
	t.PlusOne1()
	fmt.Println(t.a)
	t.PlusOne2()
	fmt.Println(t.a)
}