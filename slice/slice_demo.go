package main

import "fmt"

func main() {
	s := []int{0}
	s1 := s[:0]
	fmt.Println(len(s), cap(s), s)
	fmt.Println(len(s1), cap(s1), s1)

}
