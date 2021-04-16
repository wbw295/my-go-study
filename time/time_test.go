package main

import (
	"fmt"
	"testing"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

var BeanNotDeletedAt string= "1970-01-01 00:00:00"

func TestTime(t *testing.T) {

}

func TestTimeParse(t1 *testing.T) {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2018-09-20 15:39:06 +0800 CST")
	fmt.Println(t)
	t, _ = time.Parse("2006-01-02 15:04:05 -0700 MST", "2018-09-20 15:39:06 +0000 CST")
	fmt.Println(t)
	t, _ = time.Parse("2006-01-02 15:04:05 Z0700 MST", "2018-09-20 15:39:06 +0800 CST")
	fmt.Println(t)
	t, _ = time.Parse("2006-01-02 15:04:05 Z0700 MST", "2018-09-20 15:39:06 Z GMT")
	fmt.Println(t)
	t, _ = time.Parse("2006-01-02 15:04:05 Z0700 MST", "2018-09-20 15:39:06 +0000 GMT")
	fmt.Println(t)

	now := time.Now()
	fmt.Println(now)

	t, _ = time.Parse(TIME_LAYOUT, BeanNotDeletedAt)
	fmt.Println(t)
	fmt.Println(now.UnixNano())
	fmt.Println(int64(1 *time.Minute))
}
