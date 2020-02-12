/*
 * @Description: 类型转化
 * @Author: leo
 * @Date: 2020-02-12 14:18:26
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 14:29:03
 */
// 类型转化
package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64

// 类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// b = a 不能隐式类型转换
	b = int64(a)

	var c MyInt
	// c = b 别名隐式类型转换也不支持
	c = MyInt(b)
	fmt.Println(a, b, c)
}

// 指针
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 不支持指针运算
	fmt.Println(a, aPtr)
}

func TestString(t *testing.T) {
	var s string // 默认初始化是 空字符串
	fmt.Println("*" + s + "*")
	fmt.Println("长度为=============", len(s))
}
