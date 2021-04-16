package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

func TestChunk(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 7, 2, 5}
	s1 := funk.UniqInt(s)
	chunk := funk.Chunk(s1, 3)
	fmt.Println(chunk)

	s2 := []int{99,999}
	joinInt := funk.OuterJoinInt(s, s2)
	rightJoinInt := funk.RightJoinInt(s, s2)
	leftJoinInt := funk.LeftJoinInt(s, s2)
	fmt.Println(joinInt)
	fmt.Println(rightJoinInt)
	fmt.Println(leftJoinInt)
}

type person struct {
	Id int
	sex int
}

func TestToMap(t *testing.T)  {
	var s []*person
	s = append(s, &person{1,3})
	s = append(s, &person{2,3})
	s = append(s, &person{2,4})
	s = append(s, &person{5,7})
	m := funk.ToMap(s, "id")
	fmt.Println(m)

}
