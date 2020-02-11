package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// 简写为
	var (
		a int = 1
		b int = 1
	)
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		fmt.Println(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	// 简写方式
	a, b = b, a
	fmt.Println(a, b)
}
