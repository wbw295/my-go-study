package main

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	file, err := os.Open("path")
	if err != nil {
		return
	}

	//放在判断err状态之后
	defer file.Close()

	//todo
	//...

	return
	//defer执行时机
}
