/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 16:40:48
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 18:44:12
 */
package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回多个值的函数
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数作为参数和 函数作为返回值
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time==========spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	// time.Sleep(time.Second * 1)
	return op
}

// 函数作为参数和 函数作为返回值
func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	fmt.Println(a)
	teSF := timeSpent(slowFun)
	fmt.Println(teSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// 可变参数
func TestVarParam(t *testing.T) {
	fmt.Println(Sum(1, 2, 3, 4))
	fmt.Println(Sum(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("Clear resources")
	}()
	fmt.Println("Start")
	// panic("error")
}
