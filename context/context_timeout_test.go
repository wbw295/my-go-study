package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtxTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
		fmt.Println("main", ctx.Err().Error())
	}

	<-ctx.Done()
	fmt.Println("main2", ctx.Err())
	fmt.Println("main2", ctx.Err().Error())

	<-ctx.Done()
	fmt.Println("main3", ctx.Err())
	fmt.Println("main3", ctx.Err().Error())
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
		fmt.Println("handle", ctx.Err().Error())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}

	<-ctx.Done()
	fmt.Println("handle2", ctx.Err())
	fmt.Println("handle2", ctx.Err().Error())

	<-ctx.Done()
	fmt.Println("handle 3", ctx.Err())
	fmt.Println("handle 3", ctx.Err().Error())
}
