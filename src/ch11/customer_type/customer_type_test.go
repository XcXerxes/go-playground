/*
 * @Description: 自定义类型
 * @Author: leo
 * @Date: 2020-02-12 19:24:14
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 19:28:07
 */
package customer_type_test

import (
	"fmt"
	"testing"
	"time"
)

// 定义一个自定义类型别名
type IntConv func(op int) int

// 函数作为参数和 函数作为返回值
func timeSpent(inner IntConv) IntConv {
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
	teSF := timeSpent(slowFun)
	fmt.Println(teSF(10))
}
