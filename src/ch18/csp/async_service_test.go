package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}
func Asyncservice() chan string {
	// 普通channel
	// retCh := make(chan string)
	// 带容量的channel
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("retrune result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}
func TestService(t *testing.T) {
	retCh := Asyncservice()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
