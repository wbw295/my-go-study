package main

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {

	ins := []int{1, 2, 3, 4, 5, 6, 7, 8}

	ins = ins[2:]
	fmt.Println(ins)

}
