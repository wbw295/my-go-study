package main

import (
	"fmt"
	"testing"
)

func TestDeferError(t *testing.T){
	err := outerFunc1(1)
	if err != nil{
		fmt.Println("print err in main:",err)
	}
	fmt.Println("Done")
}

func outerFunc(param int)(err error){
	defer fmt.Println("print err in defer func:",err)
	fmt.Println("Do no thing in outerFunc")
	if param > 0{
		err = fmt.Errorf("error param:%d", param)
		return err
	}
	return nil
}

func outerFunc1(param int)(err error){
	ep := &err
	defer func(err *error) {
		fmt.Println("print err in defer func:",(*ep).Error())
	}(&err)
	fmt.Println("Do no thing in outerFunc")
	if param > 0{
		err = fmt.Errorf("error param:%d", param)
		return err
	}
	return nil
}