package main

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var v i
	defer func (v *i) {
		(*v).Handle()
	}(&v)
	v = &s{}
}

type i interface {
	Handle()
}

type s struct {

}

func (c *s) Handle(){
	fmt.Println("hello")
}