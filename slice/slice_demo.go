package main

import (
	"github.com/wbw295/my-go-study/common/log"
	"strings"
)

func main() {

	authContent := ""
	authArray := strings.Split(authContent, " ")
	for _, e := range authArray {
		log.Info(e)
	}


}
