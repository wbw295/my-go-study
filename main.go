package main

import "fmt"

type person struct {
	a int
	b int
	s *string
}

var s = "kkk"

var p = &person{1,2, &s}

func main() {
	fmt.Println(p)
	p1 := *p
	fmt.Println(*p1.s)
	*p.s = "fssfdsfds"
	fmt.Println(*p1.s)
}