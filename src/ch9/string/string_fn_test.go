/*
 * @Description: string 方法
 * @Author: leo
 * @Date: 2020-02-12 16:30:57
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 16:37:21
 */
package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"

	// 分割
	parts := strings.Split(s, ",")

	for _, part := range parts {
		fmt.Println(part)
	}

	// 字符串连接
	fmt.Println(strings.Join(parts, "-"))
}

// 字符串转换
func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	fmt.Println("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Println(10 + i)
	}
}
