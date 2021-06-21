package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello world")
	deferTest()
}

func deferTest() {
	var ii int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3*time.Second)
		fmt.Printf("go: str-> %d\n", ii)
		fmt.Printf("go: str address %p\n", &ii)
	}()
	defer func() {
		fmt.Println(ii)
		fmt.Println(&ii)
	}()
	ii = 1232342
	fmt.Println(&ii)
	wg.Wait()

}
