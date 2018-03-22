package main

import "fmt"

func main() {
	list := new([]int)
	//list := make([]int,4)
	list = append(list, 1)
	fmt.Println(list)
}
