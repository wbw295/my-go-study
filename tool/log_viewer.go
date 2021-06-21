package main

import (
	"encoding/json"
	"fmt"
	"github.com/wbw295/my-go-study/common/log"
	"io/ioutil"
)

type LogBody struct {
	Msg string `json:"msg,omitempty"`
	ErrorVerbose string `json:"errorVerbose,omitempty"`
}

func main() {

	c, err := ioutil.ReadFile("tool/log.json")
	if err != nil {
		log.Fatalp(err)
	}
	var lb LogBody
	err = json.Unmarshal(c, &lb)
	if err != nil {
		log.Fatalp(err)
	}
	fmt.Printf("message: %s\n", lb.Msg)
	fmt.Printf("verbose: %s\n", fmt.Sprintf(lb.ErrorVerbose))
}
