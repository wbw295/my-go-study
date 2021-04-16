package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)
	defer cancel()
	_, c1 := context.WithCancel(ctx)
	defer c1()
	go func() {
		time.Sleep(3*time.Second)
		c1()
		fmt.Println("gr done.")
	}()
	time.Sleep(2 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}
