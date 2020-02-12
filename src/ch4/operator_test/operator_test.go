/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 14:33:29
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 14:45:24
 */
package operator_test

import (
	"fmt"
	"testing"
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

// 数组比较
func TestCompareArray(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	var b [4]int = [4]int{1, 3, 4, 5}
	// c := [5]int{1, 2, 3, 4, 5}
	d := [4]int{1, 2, 3, 4}
	fmt.Println(a == b)
	// fmt.Println(a == c) // 长度不等 不能比较 会报错
	fmt.Println(a == d)
}

// 位运算符
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Writable
	fmt.Println("位运算===============", a&Readable == Readable, a&Writable == Writable)
}
