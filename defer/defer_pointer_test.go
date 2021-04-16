package main

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	p := &s{}
	fmt.Printf("%p\n", p)
	defer func(p *s) {
		fmt.Printf("%p\n", p)
	}(p)

}

