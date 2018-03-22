package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		println(&stu)
		fmt.Println(stu)
	}
	//for k,v:=range m{
	//	println(k,"=>",v.Name)
	//}

	for k, _ := range stus {
		m[stus[k].Name] = &stus[k]
		println(&stus[k])
		fmt.Println(stus[k])
	}
	for k,v:=range m{
		println(k,"=>",v.Name)
	}
}

